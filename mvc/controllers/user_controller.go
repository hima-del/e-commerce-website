package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../auth"
	"../model"
	"../services"
	"github.com/dgrijalva/jwt-go"
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
		result := services.Usernameexists(creds.Username)
		var s string = "username already taken"
		if result != "" {
			stringdata, err := json.Marshal(s)
			if err != nil {
				fmt.Println(err)
			}
			w.Write(stringdata)
		} else {
			token, err := services.Signup(creds.Username, creds.Password)
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
			token, err := services.Login(creds.Username, creds.Password)
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

// func CreateProduct(w http.ResponseWriter, req *http.Request) {

// }
