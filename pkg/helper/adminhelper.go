package helper

import (
	"CrocsClub/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	claims := &authCustomClaims{
		Id:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
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
