package model

import "time"

type Product struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	Picture string  `json:"picture"`
	Created *time.Time
}
