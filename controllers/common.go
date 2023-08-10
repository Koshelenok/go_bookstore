package controllers

import (
	"bookstore/services"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(services.IdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[services.IdentityKey],
		"userName": user.(*services.User).UserName,
		"text":     "Hello World.",
	})
}
