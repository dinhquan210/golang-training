package security

import (
	"golang-training/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const JWT_KEY = "hhhgfdshgfhsdgfshjgfshjdgf" //key ma hoa

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY)) //ham ma hoa token
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
