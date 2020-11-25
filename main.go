package main

import (
	"database/sql"
	"fmt"
	"net/http"

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

func main() {
	http.HandleFunc("/create", create)
	http.HandleFunc("/update", update)
	http.HandleFunc("/list", list)
	//fs := http.FileServer(http.Dir("/home/ubuntu/products"))
	//http.Handle("/products/", http.StripPrefix("/products", fs))
	http.ListenAndServe(":80", nil)
}

func create(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func update(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}

func list(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, req)
}
