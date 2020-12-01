package services

import (
	"fmt"

	"../auth"
	"../dao"
	"../model"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Usernameexists(username string) (result string) {
	result = dao.QueryOne(username)
	return
}

func Signup(username, password string) (token *model.TokenDetails, err error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	dao.QueryTwo(username, string(hashedpassword))
	id := dao.QueryThree(username)
	token, err = auth.CreateToken(id, username)
	return token, err
}

func Login(username, password string) (token *model.TokenDetails, err error) {
	storedCreds := dao.QueryFour(username)
	if err := bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(password)); err != nil {
		fmt.Println(err)
	}
	id := dao.QueryThree(username)
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
