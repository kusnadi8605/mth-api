package models

import (
	"crypto/sha256"
	"errors"
	"fmt"
	conf "mth-api/config"
	dts "mth-api/datastruct"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//GetToken is
func GetToken(conn *conf.Connection, userName string, password string) ([]dts.Token, error) {
	arrToken := []dts.Token{}
	strToken := dts.Token{}
	strUser := dts.User{}

	h := sha256.New()
	h.Write([]byte(password))

	encryptKey := fmt.Sprintf("%x", h.Sum(nil))

	strQuery := "SELECT userName,userPassword FROM mtr_user WHERE userName=? and userPassword=?"
	rows, err := conn.Query(strQuery, userName, encryptKey)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&strUser.UserName, &strUser.UserPassword)
		if err != nil {
			return nil, err
		}
	}

	if strUser.UserName != userName || encryptKey != strUser.UserPassword {
		fmt.Println("invalid userName or password")
		return nil, fmt.Errorf("invalid userName or password")
	}
	//generate Token
	token, _ := GenToken(conf.Param.JwtKEY)

	strToken.Token = token
	arrToken = append(arrToken, strToken)

	return arrToken, nil
}

//GenToken ..
func GenToken(key string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	var mySigningKey = []byte(key)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	/* Sign the token with our secret */
	tokenString, err := token.SignedString(mySigningKey)
	return tokenString, err
}

//ValidToken ..
func ValidToken(r *http.Request) (bool, error) {
	var mySigningKey = []byte(conf.Param.JwtKEY)
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, errors.New("Invalid Token")
	}

	return true, nil
}
