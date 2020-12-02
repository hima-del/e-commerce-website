package services

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"strings"

	"../dao"
	"../model"
)

func CreateProduct(n string, s float64, file multipart.File) (pd model.Product, err error) {
	tempFile, err := ioutil.TempFile("products", "upload-*.png")
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	tempFile.Write(fileBytes)
	fmt.Println("successfully uploaded file")
	fileName := tempFile.Name()
	v := strings.TrimPrefix(fileName, `products\`)
	err = dao.QueryFive(n, s, v)
	id, err := dao.QuerySix(v)
	pd, err = dao.QuerySeven(id)
	return pd, err
}

func GetProducts() (productslist []model.Product, err error) {
	productList, err := dao.QueryEight()
	return productList, err
}

func GetSingleProduct(id int) (product model.Product, err error) {
	product, err = dao.QuerySeven(id)
	return product, err
}

func DeleteProduct(id int) (err error) {
	deletedpicture, err := dao.QueryNine(id)
	picturename := "products/" + deletedpicture
	err = os.Remove(picturename)
	err = dao.QueryTen(id)
	return err
}

func UpdateProduct(id int, n string, p string, file multipart.File) (product model.Product, err error) {
	pd, err := dao.QuerySeven(id)
	if p != "" {
		s, err := strconv.ParseFloat(p, 64)
		err = dao.QueryTwelve(s, id)
		fmt.Println(err)
	}
	if n != "" {
		err = dao.QueryThirteen(n, id)
	}

	if file != nil {
		tempFile, err := ioutil.TempFile("products", "upload-*.png")
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		tempFile.Write(fileBytes)
		fmt.Println("successfully uploaded file")
		fileName := tempFile.Name()
		fmt.Println(fileName)
		v := strings.TrimPrefix(fileName, `products\`)
		err = dao.QueryForteen(v, id)
		picturename := "products/" + pd.Picture
		err = os.Remove(picturename)
		fmt.Println(err)
	}
	product, err = dao.QuerySeven(id)
	return product, err
}

func Logout(tokenStringLogout string) (err error) {
	err = dao.QueryEleven(tokenStringLogout)
	return err
}
