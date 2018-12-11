package main

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func generateToken(email string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
	})
	tokenString, error := token.SignedString([]byte(os.Getenv("RSAKey")))
	if error != nil {
		fmt.Println(error)
	}
	return tokenString
}
