package model

import "github.com/dgrijalva/jwt-go"

type JwtCustomClaims struct {
	UserId   string
	Role     string
	FullName string
	jwt.StandardClaims
}
