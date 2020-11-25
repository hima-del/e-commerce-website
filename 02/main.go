package main

import (
	"github.com/gorilla/mux"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://himaja:password@localhost/ecommerce?sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you are connected to database")
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func main(){
	r:=mux.NewRouter()
	r.HandleFunc("/signup",signup)
	r.HandleFunc("/login",login)
	r.HandleFunc("/products/"getProduct)
	r.HandleFunc("/products/",createProduct)
	r.HandleFunc("/products/:id"updateProduct)
	r.HandleFunc("/products/:id",getSingleProduct)
	r.HandleFunc("/products/:id"deleteProduct)
	r.HandleFunc("/logout",logout)
}
