package services

import (
	"fmt"

	"../dao"
	"github.com/dgrijalva/jwt-go"
)

func CreateOrder(ID, quantity int, orderDate, shippingDate, orderStatus, billingAddress, shippingAddress, size, color string, price, discount, total float64) (err error) {
	err = dao.QueryFifteen(ID, billingAddress, shippingAddress)
	err = dao.QuerySixteen(ID, orderDate, shippingDate, shippingAddress, orderStatus)
	orderID, err := dao.QuerySeventeen(ID)
	err = dao.QueryEighteen(orderID, price, discount, total, quantity, color, size)
	return err
}

func ExtractID(tokenString string) (ID int, err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("ACCESS_SECRET"), nil
	})
	fmt.Println(token)
	extractedID := claims["user_id"]
	v, ok := extractedID.(float64)
	fmt.Println(v, ok)
	ID = int(v)
	return ID, err
}

func assert(i interface{}) {
	v, ok := i.(float64)
	fmt.Println(v, ok)
}
