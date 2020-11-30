package dao

import (
	"fmt"

	"../model"
)

func QueryOne(username string) (storedCreds *model.Credentials) {
	result := config.db.QueryRow("select username from customers where username=$1", username)
	storedCreds = &model.Credentials{}
	err := result.Scan(&storedCreds.Username)
	if err != nil {
		fmt.Println(err)
	}
	return storedCreds
}

func QueryTwo(username, password string) {
	_, err := config.db.Query("insert into customers (username,password)values ($1,$2)", username, password)
	if err != nil {
		fmt.Println(err)
	}
}

func QueryThree(username string) (id int) {
	resultID := config.db.QueryRow("select id from customers where username=$1", username)
	err := resultID.Scan(&id)
	return id
}
