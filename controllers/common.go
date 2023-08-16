package controllers

import (
	"bookstore/models"
	"bookstore/services"
	"fmt"
	"strconv"

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

func getIDParam(c *gin.Context) (id int, err error) {
	rawID := c.Param("id")
	parsedID, err := strconv.Atoi(rawID)
	if err != nil {
		return 0, fmt.Errorf("incorrect ID parameter")
	}
	return parsedID, nil
}
