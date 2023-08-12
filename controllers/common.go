package controllers

import (
	"bookstore/models"
	"bookstore/services"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(services.AuthIdentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[services.AuthIdentityKey],
		"userName": user.(*models.User).FirstName,
		"text":     "Hello World.",
	})
}
