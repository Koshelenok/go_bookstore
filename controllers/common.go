package controllers

import (
	"bookstore/models"
	"bookstore/services"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(models.UserIdkentityKey)
	c.JSON(200, gin.H{
		"userID":   claims[models.UserIdkentityKey],
		"userName": user.(*services.LoggedUser).Name,
		"text":     "Hello World.",
	})
}
