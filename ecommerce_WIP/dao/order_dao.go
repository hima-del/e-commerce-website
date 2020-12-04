package dao

import (
	"../config"
)

func QueryFifteen(ID int, billingAddress, shippingAddress string) (err error) {
	_, err = config.DB.Query("insert into address (customerid,billingaddress,shippingaddress)values ($1,$2,$3)", ID, billingAddress, shippingAddress)
	return err
}

func QuerySixteen(ID int, orderDate, shippingDate, shippingAddress, orderStatus string) (err error) {
	_, err = config.DB.Query("insert into orders (customerid,orderdate,shipdate,shippingaddress, orderstatus)values ($1,$2,$3,$4,$5)", ID, orderDate, shippingDate, shippingAddress, orderStatus)
	return err
}

func QuerySeventeen(ID int) (id int, err error) {
	result := config.DB.QueryRow("select id from orders where customerid=$1", ID)
	err = result.Scan(&id)
	return id, err
}

func QueryEighteen(productid, orderID, ordernumber int, price, discount, total float64, quantity int, color, size string) (err error) {
	_, err = config.DB.Query("insert into orderdetails (productid,orderid,ordernumber, price, discount, total,quantity,color,size)values ($1,$2,$3,$4,$5,$6,$7,$8,$9)", productid, orderID, ordernumber, price, discount, total, quantity, color, size)
	return err
}
