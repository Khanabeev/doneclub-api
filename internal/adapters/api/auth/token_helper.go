package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type accessTokenClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func ClaimsFromToken(tokenString string) (*accessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &accessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		fmt.Println("Error while parsing token: " + err.Error())
		return nil, err
	}
	claims := token.Claims.(*accessTokenClaims)
	return claims, nil
}
