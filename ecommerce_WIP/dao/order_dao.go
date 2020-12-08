package dao

import (
	"../config"
	"../model"
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

func QueryNineteen(ordernumber int) (id int, err error) {
	result := config.DB.QueryRow("select id from orderdetails where ordernumber=$1", ordernumber)
	err = result.Scan(&id)
	return id, err
}

func QueryTwenty(id int) (order model.OrderDetails, err error) {
	row := config.DB.QueryRow("select * from orderdetails where id=$1", id)
	order = model.OrderDetails{}
	err = row.Scan(&order.ID, &order.Productid, &order.Orderid, &order.Ordernumber, &order.Price, &order.Discount, &order.Total, &order.Quantity, &order.Color, &order.Size, &order.Created)
	return order, err
}

func QueryTwentyone(id int) (err error) {
	_, err = config.DB.Query("delete from orderdetails where id=$1", id)
	return err
}
