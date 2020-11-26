package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://himaja:password@localhost/ecommerce_web?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you are connected to database")
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	ATExpires    int64
	RTExpires    int64
}

type Product struct {
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Picture string  `json:"picture"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/login", login)
	r.HandleFunc("/products/", getProducts).Methods("GET")
	r.HandleFunc("/products/", createProduct).Methods("POST")
	//r.HandleFunc("/products/:id", updateProduct)
	r.HandleFunc("/products/:id", getSingleProduct).Methods("GET")
	//r.HandleFunc("/products/:id", deleteProduct)
	r.HandleFunc("/logout", logout)
	http.ListenAndServe(":80", r)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if req.URL.Path != "/signup" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var creds Credentials
		err = json.Unmarshal(b, &creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := db.QueryRow("select username from customers where username=$1", creds.Username)
		storedCreds := &Credentials{}
		err = result.Scan(&storedCreds.Username)
		var s string = "username already taken"
		if storedCreds.Username != "" {
			stringdata, err := json.Marshal(s)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(stringdata)
		} else {
			hashedpassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
			_, err = db.Query("insert into customers (username,password)values ($1,$2)", creds.Username, string(hashedpassword))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			resultID := db.QueryRow("select id from customers where username=$1", creds.Username)
			var id int
			err = resultID.Scan(&id)
			//fmt.Println(id)
			token, err := createToken(id, creds.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			tokens := map[string]string{
				"acces_token":   token.AccessToken,
				"refresh_token": token.RefreshToken,
			}
			data, err := json.Marshal(tokens)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(data)
		}
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if req.URL.Path != "/login" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		beartoken := req.Header.Get("Authorization")
		if beartoken == "" {
			creds := &Credentials{}
			err := json.NewDecoder(req.Body).Decode(creds)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			result := db.QueryRow("select password from customers where username=$1", creds.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			storedCreds := &Credentials{}
			err = result.Scan(&storedCreds.Password)
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
			}
			//if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err == nil {
			resultID := db.QueryRow("select id from customers where username=$1", creds.Username)
			fmt.Println(resultID)
			var id int
			err = resultID.Scan(&id)
			//fmt.Println("id", id)
			token, err := createToken(id, creds.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Println("access token length", len(token.AccessToken))
			fmt.Println("refresh token length", len(token.RefreshToken))
			tokens := map[string]string{
				"acces_token":   token.AccessToken,
				"refresh_token": token.RefreshToken,
			}
			data, err := json.Marshal(tokens)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(data)
		} else if beartoken != "" {
			//fmt.Println("entered")
			tokenString := extractToken(req)
			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte("REFRESH_SECRET"), nil
			})
			fmt.Println(token)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// ... error handling

			// do something with decoded claims
			for key, val := range claims {
				fmt.Printf("Key: %v, value: %v\n", key, val)
			}
			_, ok := claims["refresh_uuid"]
			if ok == true {
				idExtracted := claims["username"]
				//fmt.Println(id)
				newAccesstoken, err := createAccessToken(idExtracted)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				newToken := map[string]string{
					"access_token": newAccesstoken.AccessToken,
				}
				tokenData, err := json.Marshal(newToken)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				w.Write(tokenData)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}
	}
}

func getSingleProduct(w http.ResponseWriter, req *http.Request) {
	fmt.Println("bye")
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	fmt.Println("hello")
	blacklistToken := checkBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := tokenValid(w, req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println("hi")
	vars := mux.Vars(req)
	fmt.Println("vars", vars)
}

func getProducts(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	blacklistToken := checkBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := tokenValid(w, req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	rows, err := db.Query("select * from products")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	productslist := make([]Product, 0)
	for rows.Next() {
		pdct := Product{}
		err := rows.Scan(&pdct.Name, &pdct.Price, &pdct.Picture)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		productslist = append(productslist, pdct)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	js, err := json.Marshal(productslist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	req.ParseMultipartForm(10 << 20)
	blacklistToken := checkBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := tokenValid(w, req)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	n := req.FormValue("name")
	p := req.FormValue("price")
	s, err := strconv.ParseFloat(p, 64)
	fmt.Printf("%T", s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//file uploading
	file, header, err := req.FormFile("picture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("uploaded file:", header.Filename)
	fmt.Println("file size:", header.Size)
	fmt.Println("MIME header:", header.Header)
	tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tempFile.Write(fileBytes)
	fmt.Println("successfully uploaded file")
	fileName := tempFile.Name()
	fmt.Println(fileName)
	v := strings.TrimPrefix(fileName, `temp-images\`)
	//pdcts := Product{n, s, v}
	//fmt.Println("pdcts", pdcts)
	_, err = db.Query("insert into products (name,price,picture)values($1,$2,$3)", n, s, v)
	if err != nil {
		fmt.Println("internal server error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resultCreate := db.QueryRow("select * from products where name=$1", n)
	pd := Product{}
	resultCreate.Scan(&pd.Name, &pd.Price, &pd.Picture)
	fmt.Println("pd", pd)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func logout(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		blacklistToken := checkBlacklist(w, req)
		if blacklistToken != "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			err := tokenValid(w, req)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			tokenStringLogout := extractToken(req)
			stmnt := "insert into blacklist (token)values ($1)"
			_, err = db.Exec(stmnt, tokenStringLogout)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println("succesfully logged out")
		}
	}
}
