package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
			returns
		}
		var creds model.Credentials
		err = json.Unmarshal(b, &creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
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

func Login(w http.ResponseWriter, req *http.Request) {
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
			token, err := services.Login(creds.Username)
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
			_, ok := claims["refresh_uuid"]
			if ok == true {
				idExtracted := claims["username"]
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
