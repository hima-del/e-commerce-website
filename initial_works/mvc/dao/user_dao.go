package dao

import (
	"../config"
)

func QueryOne(username string) (name string, err error) {
	result := config.DB.QueryRow("select username from customers where username=$1", username)
	err = result.Scan(&name)
	return name, err
}

func QueryTwo(username, password string) (err error) {
	_, err = config.DB.Query("insert into customers (username,password)values ($1,$2)", username, password)
	return err
}

func QueryThree(username string) (id int, err error) {
	resultID := config.DB.QueryRow("select id from customers where username=$1", username)
	err = resultID.Scan(&id)
	return id, err
}

func QueryFour(username string) (storedPassword string, err error) {
	result := config.DB.QueryRow("select password from customers where username=$1", username)
	err = result.Scan(&storedPassword)
	return storedPassword, err
}
