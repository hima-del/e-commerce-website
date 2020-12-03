package services

import (
	"fmt"

	"../dao"
	"github.com/dgrijalva/jwt-go"
)

func CreateOrder(tokenString, orderDate, shippingDate, orderStatus, billingAddress, shippingAddress string, price, quantity, discount, total float64, size, color string) (err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("ACCESS_SECRET"), nil
	})
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
	}
	extrctedID := claims["user_id"]
	err = dao.QueryFifteen(extrctedID, billingAddress, shippingAddress)
	err = dao.QuerySixteen(extrctedID, orderDate, shippingDate, shippingAddress, orderStatus)
	err = dao.QuerySeventeen()
}
