package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"../auth"
	"../services"
	"github.com/gorilla/mux"
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
	n := req.FormValue("name")
	p := req.FormValue("price")
	s, err := strconv.ParseFloat(p, 64)
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

func GetProducts(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	blacklistToken := auth.CheckBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	productsList, err := services.GetProducts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(productsList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetSingleProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	blacklistToken := auth.CheckBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	product, err := services.GetSingleProduct(id)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	blacklistToken := auth.CheckBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = services.DeleteProduct(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateProduct(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
	req.ParseMultipartForm(10 << 20)
	blacklistToken := auth.CheckBlacklist(w, req)
	if blacklistToken != "" {
		w.WriteHeader(http.StatusUnauthorized)
	}
	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	n := req.PostFormValue("name")
	p := req.PostFormValue("price")
	file, header, err := req.FormFile("picture")
	defer file.Close()
	fmt.Println("uploaded file:", header.Filename)
	fmt.Println("file size:", header.Size)
	fmt.Println("MIME header:", header.Header)
	product, err := services.UpdateProduct(id, n, p, file)
	switch {
	case err == sql.ErrNoRows:
		http.NotFound(w, req)
		return
	case err != nil:
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func Logout(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		blacklistToken := auth.CheckBlacklist(w, req)
		if blacklistToken != "" {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			tokenStringLogout := auth.ExtractToken(req)
			err := services.Logout(tokenStringLogout)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Println("successfully logged out")
		}
	}
}
