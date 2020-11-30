package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../model"
	"../services"
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


func Login