package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	secret := os.Getenv("LETSPLAY_JWT_SECRET")
	if secret == "" {
		panic("LETSPLAY_JWT_SECRET undefined")
	}

	claims := jwt.MapClaims{
		"role": "service_role",
		"sub":  "admin",
		"exp":  time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
}
