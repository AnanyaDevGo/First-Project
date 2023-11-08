package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AdminAuthMiddleware(c *gin.Context) {

	fmt.Println("**** admin auth*****")
	accessToken := c.Request.Header.Get("Authorization")

	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	_, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte("accesssecret"), nil
	})
	fmt.Println("here  1")
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	c.Next()
}
