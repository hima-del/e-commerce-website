package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../auth"
	"../model"
	"../services"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, req *http.Request) {
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
		var creds model.Credentials
		err = json.Unmarshal(b, &creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result, err := services.Usernameexists(creds.Username)
		var s string = "username already taken"
		if result != "" {
			stringdata, err := json.Marshal(s)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(stringdata)
		} else {
			hashedPassword, err := services.HashPassword(creds.Password)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = services.InsertToCustomers(creds.Username, string(hashedPassword))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			token, err := services.Signup(creds.Username)
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

func Login(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		if req.URL.Path != "/login" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		beartoken := req.Header.Get("Authorization")
		if beartoken == "" {
			creds := &model.Credentials{}
			err := json.NewDecoder(req.Body).Decode(creds)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			storedPassword, err := services.ExistingUser(creds.Username)
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(creds.Password))
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token, err := services.Login(creds.Username)
			if err != nil {
				if err == sql.ErrNoRows {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
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
		} else if beartoken != "" {
			claims := jwt.MapClaims{}
			tokenString := auth.ExtractToken(req)
			ok, err := services.Refresh(tokenString)
			if err != nil {
				fmt.Println(err)
			}
			if ok == true {
				idExtracted := claims["username"]
				newAccesstoken, err := auth.CreateAccessToken(idExtracted)
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
