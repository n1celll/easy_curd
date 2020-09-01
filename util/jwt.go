package util

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

const (
	TokenIssue          = "SERVER"
	TokenExpireDuration = time.Second * 60 * 60 * 24 * 7
	//TokenExpireDuration = time.Second * 3

)

var secret = os.Getenv("SERVER_SECRET")

func GenJwt(uid uint) string {
	_token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		Issuer:    TokenIssue,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		Subject:   strconv.Itoa(int(uid)),
	})
	token, _ := _token.SignedString([]byte(secret))
	return token
}

func DecodeJwt(tokenString string) (*jwt.StandardClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return claims, err
	} else {
		return nil, err
	}
}
