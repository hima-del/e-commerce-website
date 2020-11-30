package services

import (
	"encoding/json"
	"fmt"

	"../dao"
	"golang.org/x/crypto/bcrypt"
)

func Signup(username string) (*model.TokenDetails, error) {
	result := dao.QueryOne(username)
	var s string = "username already taken"
	if result != "" {
		stringdata, err := json.Marshal(s)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(stringdata)
	} else {
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), 8)
		dao.QueryTwo(creds.Username, string(hashedpassword))
		id := dao.QueryThree(creds.Username)
		token, err := jwt.CreateToken(id, creds.Username)
		if err != nil {
			fmt.Println(err)
		}
		return token, err
	}
}

func Login(username string){
	storedCreds:=dao.QueryFour(username)
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	}
	id:=dao.QueryThree(username)
	token, err := jwt.CreateToken(id, creds.Username)
		if err != nil {
			fmt.Println(err)
		}
		return token, err
	}
}
