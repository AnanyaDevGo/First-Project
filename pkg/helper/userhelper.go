package helper

import (
	"CrocsClub/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
	claims := &authCustomClaims{
		Id:    user.Id,
		Email: user.Email,
		Role:  "client",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("comebuycrocs"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
