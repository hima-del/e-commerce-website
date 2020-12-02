package services

import (
	"fmt"

	"../auth"
	"../dao"
	"../model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Usernameexists(username string) (result string, err error) {
	result, err = dao.QueryOne(username)
	return result, err
}

func HashPassword(password string) (hashedpassword []byte, err error) {
	hashedpassword, err = bcrypt.GenerateFromPassword([]byte(password), 8)
	return hashedpassword, err
}

func InsertToCustomers(username, password string) (err error) {
	err = dao.QueryTwo(username, password)
	return err
}

func Signup(username string) (token *model.TokenDetails, err error) {
	id, err := dao.QueryThree(username)
	token, err = auth.CreateToken(id, username)
	return token, err
}

func ExistingUser(username string) (storedPassword string, err error) {
	storedPassword, err = dao.QueryFour(username)
	return storedPassword, err
}

func Login(username string) (token *model.TokenDetails, err error) {
	id, err := dao.QueryThree(username)
	token, err = auth.CreateToken(id, username)
	return token, err
}

func Refresh(tokenString string) (ok bool, err error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("REFRESH_SECRET"), nil
	})
	fmt.Println(token)
	_, ok = claims["refresh_uuid"]
	return ok, err
}
