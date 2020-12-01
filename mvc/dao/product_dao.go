package dao

import (
	"fmt"

	"../config"
	"../model"
)

func QueryFive(n string, s float64, v string) {
	_, err := config.DB.Query("insert into products (name,price,picture)values($1,$2,$3)", n, s, v)
	if err != nil {
		fmt.Println(err)
	}
}

func QuerySix(v string) (id int) {
	idRow := config.DB.QueryRow("select id from products where picture=$1", v)
	err := idRow.Scan(&id)
	fmt.Println("id", id)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func QuerySeven(id int) (pd *model.Product, err error) {
	row := config.DB.QueryRow("select * from products where id=$1", id)
	pd = &model.Product{}
	err = row.Scan(&pd.ID, &pd.Name, &pd.Price, &pd.Picture, &pd.Created)
	return pd, err
}
