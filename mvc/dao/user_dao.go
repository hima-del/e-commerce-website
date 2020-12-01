package dao

import (
	"database/sql"
	"fmt"

	"../config"
	"../model"
)

func QueryOne(username string) (name string) {
	result := config.DB.QueryRow("select username from customers where username=$1", username)
	err := result.Scan(&name)
	if err != nil {
		fmt.Println(err)
	}
	return name
}

func QueryTwo(username, password string) {
	_, err := config.DB.Query("insert into customers (username,password)values ($1,$2)", username, password)
	if err != nil {
		fmt.Println(err)
	}
}

func QueryThree(username string) (id int) {
	resultID := config.DB.QueryRow("select id from customers where username=$1", username)
	err := resultID.Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func QueryFour(username string) (storedCreds *model.Credentials) {
	result := config.DB.QueryRow("select password from customers where username=$1", username)
	storedCreds = &model.Credentials{}
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println(err)
		}
		fmt.Println(err)
	}
	return storedCreds
}
