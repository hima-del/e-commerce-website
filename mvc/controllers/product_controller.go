package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../auth"
	"../services"
)

func CreateProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	req.ParseMultipartForm(10 << 20)
	blacklistToken := auth.CheckBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err := auth.TokenValid(w, req)
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
	file, header, err := req.FormFile("picture")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	fmt.Println("uploaded file:", header.Filename)
	fmt.Println("file size:", header.Size)
	fmt.Println("MIME header:", header.Header)
	pd, err := services.CreateProduct(n, s, file)
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
