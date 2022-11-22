package jwthandle

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/praveenmahasena647/vue-link-tree/dbs"
)

func CreateJwt(person *dbs.Login) (string, error) {

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": person.Id.Hex(),
		"exp": time.Now().Add(time.Hour * 24 * 15).Unix(),
	})

	// token Key is here just for test perpose
	var tokenString, tokenErr = token.SignedString([]byte("yayayaa"))

	return tokenString, tokenErr
}

func DecodeJwt(r *http.Request) (string, error) {
	var JwtString = r.Header.Get("X-Jwt")

	var strResult string
	if JwtString == "" {
		return strResult, errors.New("no Jwt String")
	}
	var token, err = jwt.Parse(JwtString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", errors.New("its a Eoorror")
		}
		// token Key is here just for test perpose
		return []byte("yayayaa"), nil
	})
	//
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) || claims["exp"] == nil {
			return strResult, errors.New("its A Err ")
		}
		strResult = claims["_id"].(string)
	}
	return strResult, err
}

func CreateAdminJWT(key string) (string, error) {

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": key,
	})

	// token Key is here just for test perpose
	var tokenString, tokenErr = token.SignedString([]byte("yayayaa"))

	return tokenString, tokenErr
}
