package services

import (
	"fmt"

	"../dao"
	"../model"
	"github.com/dgrijalva/jwt-go"
)

func CreateOrder(ID, quantity, productid, ordernumber int, orderDate, shippingDate, orderStatus, billingAddress, shippingAddress, size, color string, price, discount, total float64) (order model.OrderDetails, err error) {
	err = dao.QueryFifteen(ID, billingAddress, shippingAddress)
	err = dao.QuerySixteen(ID, orderDate, shippingDate, shippingAddress, orderStatus)
	orderID, err := dao.QuerySeventeen(ID)
	err = dao.QueryEighteen(productid, orderID, ordernumber, price, discount, total, quantity, color, size)
	Queryid, err := dao.QueryNineteen(ordernumber)
	order, err = dao.QueryTwenty(Queryid)
	return order, err
}

func ExtractID(tokenString string) (ID int) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("ACCESS_SECRET"), nil
	})
	fmt.Println(token)
	if err != nil {
		fmt.Println(err)
	}
	extractedID := claims["user_id"]
	v, ok := extractedID.(float64)
	fmt.Println(v, ok)
	ID = int(v)
	return ID
}

func DeleteOrder(id int) (err error) {
	err = dao.QueryTwentyone(id)
	return err
}

func GetOrders() (orderList []model.OrderDetails, err error) {
	orderList, err = dao.QueryTwentytwo()
	return orderList, err
}

func GetSingleOrder(id int) (order model.OrderDetails, err error) {
	order, err = dao.QueryTwentythree(id)
	return order, err
}
