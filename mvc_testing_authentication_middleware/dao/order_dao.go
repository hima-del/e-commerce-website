package dao

import "../config"

func QueryFifteen(extrctedID int, billingAddress, shippingAddress string) (err error) {
	_, err = config.DB.Query("insert into address (customer id,billing address,shipping address)values ($1,$2,$3)", extrctedID, billingAddress, shippingAddress)
	return err
}

func QuerySixteen(extrctedID int, orderDate, shippingDate, shippingAddress, orderStatus string) (err error) {
	_, err = config.DB.Query("insert into orders (customer id,order date,ship date,shipping address, order status)values ($1,$2,$3)", extrctedID, orderDate, shippingDate, shippingAddress, orderStatus)
	return err
}
