package model

import "time"

type Order struct {
	ID              int
	Customerid      int
	Orderdate       string
	Shipdate        string
	Shippingaddress string
	Orderstatus     string
	created         *time.Time
}
