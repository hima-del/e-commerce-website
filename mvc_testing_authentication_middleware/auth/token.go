package auth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"../config"
	"../model"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

func ExtractToken(req *http.Request) string {
	bearToken := req.Header.Get("Authorization")
	stringArray := strings.Split(bearToken, " ")
	if len(stringArray) == 2 {
		return stringArray[1]
	}
	return ""
}

func CheckBlacklist(w http.ResponseWriter, req *http.Request) string {
	tokenString := ExtractToken(req)
	result := config.DB.QueryRow("select * from blacklist where token=$1", tokenString)
	var blacklistedToken string
	err := result.Scan(&blacklistedToken)
	if err != nil {
		return ""
	}
	return blacklistedToken
}

func CreateToken(userid int, username string) (*model.TokenDetails, error) {
	var err error
	td := &model.TokenDetails{}
	td.ATExpires = time.Now().Add(time.Hour * 24).Unix()
	td.RTExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	AccessID, _ := uuid.NewV4()
	td.AccessUUID = AccessID.String()
	RefreshID, _ := uuid.NewV4()
	td.RefreshUUID = RefreshID.String()

	//creating access token
	os.Setenv("ACCESS_SECRET", "accesskey")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userid
	atClaims["username"] = username
	atClaims["exp"] = td.ATExpires
	pointerToAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = pointerToAccessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//creating refresh token
	os.Setenv("REFRESH_SECRET", "refreshkey")
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUUID
	rtClaims["user_id"] = userid
	rtClaims["username"] = username
	rtClaims["exp"] = td.RTExpires
	pointerToRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = pointerToRefreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func CreateAccessToken(username interface{}) (*model.TokenDetails, error) {
	var err error
	at := &model.TokenDetails{}
	at.ATExpires = time.Now().Add(time.Hour * 24).Unix()
	AccessID, _ := uuid.NewV4()
	at.AccessUUID = AccessID.String()
	//creating new access token
	os.Setenv("ACCESS_SECRET", "newaccesssecret")
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = at.AccessUUID
	atClaims["username"] = username
	atClaims["exp"] = at.ATExpires
	pointerToAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	at.AccessToken, err = pointerToAccessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	return at, nil
}
