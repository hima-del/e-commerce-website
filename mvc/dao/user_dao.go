package dao

import (
	"database/sql"
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
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func QueryFour(username string) (storedCreds *model.Credentials) {
	result := congig.db.QueryRow("select password from customers where username=$1", creds.Username)
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
