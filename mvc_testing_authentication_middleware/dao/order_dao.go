package dao

import "../config"

func QueryFifteen(extrctedID int, billingAddress, shippingAddress string) (err error) {
	_, err = config.DB.Query("insert into address (customer id,billing address,shipping address)values ($1,$2,$3)", extrctedID, billingAddress, shippingAddress)
	return err
}

func QuerySixteen(extrctedID int, orderDate, shippingDate, shippingAddress, orderStatus string) (err error) {
	_, err = config.DB.Query("insert into orders (customer id,order date,ship date,shipping address, order status)values ($1,$2,$3,$4,$5)", extrctedID, orderDate, shippingDate, shippingAddress, orderStatus)
	return err
}

func QuerySeventeen(extractedID int) (id int, err error) {
	result := config.DB.QueryRow("select id from orders where customer id=$1", extractedID)
	err = result.Scan(&id)
	return id, err
}

func QueryEighteen(orderID int, price, discount, total float64, quantity int, color, size string) (err error) {
	_, err = config.DB.Query("insert into order_details (order id, price, discount, total,quantity,color,size)values ($1,$2,$3,$4,$5,$6,$7)", orderID, price, discount, total, quantity, color, size)
	return err
}
