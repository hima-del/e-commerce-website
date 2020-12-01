package services

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"strings"

	"../dao"
	"../model"
)

func CreateProduct(n string, s float64, file multipart.File) (pd *model.Product, err error) {
	tempFile, err := ioutil.TempFile("products", "upload-*.png")
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	tempFile.Write(fileBytes)
	fmt.Println("successfully uploaded file")
	fileName := tempFile.Name()
	v := strings.TrimPrefix(fileName, `products\`)
	dao.QueryFive(n, s, v)
	id := dao.QuerySix(v)
	pd, err = dao.QuerySeven(id)
	return pd, err
}
