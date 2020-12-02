package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://himaja:password@localhost/ecommerce?sslmode=disable")
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

func main() {
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/list", list)
	http.HandleFunc("/logout", logout)
	//fs := http.FileServer(http.Dir("/home/ubuntu/products"))
	//http.Handle("/products/", http.StripPrefix("/products", fs))
	http.ListenAndServe(":80", nil)
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
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		var creds Credentials
		err = json.Unmarshal(b, &creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result := db.QueryRow("select username from userdetails where username=$1", creds.Username)
		var storedUsername string
		err = result.Scan(&storedUsername)
		var s string = "username already taken"
		if storedUsername != "" {
			stringData, err := json.Marshal(s)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(stringData)
		} else {
			hashedpassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
			_, err = db.Query("insert into userdetails (username,password)values ($1,$2)", creds.Username, string(hashedpassword))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			//processing
		}
	}
}

func login(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func create(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func update(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func list(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func logout(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}
