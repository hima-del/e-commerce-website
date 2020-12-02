package dao

import (
	"../config"
	"../model"
)

func QueryFive(n string, s float64, v string) (err error) {
	_, err = config.DB.Query("insert into products (name,price,picture)values($1,$2,$3)", n, s, v)
	return err
}

func QuerySix(v string) (id int, err error) {
	idRow := config.DB.QueryRow("select id from products where picture=$1", v)
	err = idRow.Scan(&id)
	return id, err
}

func QuerySeven(id int) (pd model.Product, err error) {
	row := config.DB.QueryRow("select * from products where id=$1", id)
	pd = model.Product{}
	err = row.Scan(&pd.ID, &pd.Name, &pd.Price, &pd.Picture, &pd.Created)
	return pd, err
}

func QueryEight() (productslist []model.Product, err error) {
	rows, err := config.DB.Query("select * from products")
	defer rows.Close()
	productslist = make([]model.Product, 0)
	for rows.Next() {
		pdct := model.Product{}
		err = rows.Scan(&pdct.ID, &pdct.Name, &pdct.Price, &pdct.Picture, &pdct.Created)
		productslist = append(productslist, pdct)
	}
	return productslist, err
}

func QueryNine(id int) (deletedpicture string, err error) {
	result := config.DB.QueryRow("select picture from products where id=$1", id)
	err = result.Scan(&deletedpicture)
	return deletedpicture, err
}

func QueryTen(id int) (err error) {
	_, err = config.DB.Query("delete from products where id=$1", id)
	return err
}

func QueryEleven(tokenString string) (err error) {
	stmnt := "insert into blacklist (token)values ($1)"
	_, err = config.DB.Exec(stmnt, tokenString)
	return err
}

func QueryTwelve(s float64, id int) (err error) {
	_, err = config.DB.Query("update products set price=$1 where id=$2", s, id)
	return err
}

func QueryThirteen(n string, id int) (err error) {
	_, err = config.DB.Query("update products set name=$1 where id=$2", n, id)
	return err
}

func QueryForteen(v string, id int) (err error) {
	_, err = config.DB.Query("update products set picture=$1 where id=$2", v, id)
	return err
}
