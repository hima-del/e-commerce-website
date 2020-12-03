package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"../auth"
	"../services"
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
	extractedID, err := services.ExtractID(tokenString)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("id", extractedID)
	err = services.CreateOrder(extractedID, quantity, orderDate, shippingDate, orderStatus, billingAddress, shippingAddress, size, color, price, discount, total)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
