package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../auth"
	"../services"
	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, req *http.Request) {
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
	productidvalue := req.PostFormValue("product id")
	productid, err := strconv.Atoi(productidvalue)
	ordernumbervalue := req.PostFormValue("order number")
	ordernumber, err := strconv.Atoi(ordernumbervalue)
	orderDate := req.PostFormValue("order date")
	shippingDate := req.PostFormValue("shipping date")
	orderStatus := req.PostFormValue("order status")
	billingAddress := req.PostFormValue("billing address")
	shippingAddress := req.PostFormValue("shipping address")
	priceValue := req.PostFormValue("price")
	price, err := strconv.ParseFloat(priceValue, 64)
	quantityValue := req.PostFormValue("quantity")
	quantity, err := strconv.Atoi(quantityValue)
	discountValue := req.PostFormValue("discount")
	discount, err := strconv.ParseFloat(discountValue, 64)
	totalValue := req.PostFormValue("total")
	total, err := strconv.ParseFloat(totalValue, 64)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	size := req.PostFormValue("size")
	color := req.PostFormValue("color")
	tokenString := auth.ExtractToken(req)
	extractedID := services.ExtractID(tokenString)
	order, err := services.CreateOrder(extractedID, quantity, productid, ordernumber, orderDate, shippingDate, orderStatus, billingAddress, shippingAddress, size, color, price, discount, total)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	js, err := json.Marshal(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func DeleteOrder(w http.ResponseWriter, req *http.Request) {
	if req.Method != "DELETE" {
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
	err = services.DeleteOrder(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func GetOrders(w http.ResponseWriter, req *http.Request) {

}

func GetSingleOrder(w http.ResponseWriter, req *http.Request) {

}
