package model

import "time"

type OrderDetails struct {
	ID          int
	Productid   int
	Orderid     int
	Ordernumber int
	Price       float64
	Discount    float64
	Total       float64
	Quantity    int
	Color       string
	Size        string
	Created     *time.Time
}
