package controllers

import (
	"bookstore/models"
	"bookstore/services"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type UserLoginInput struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) (interface{}, error) {
	var user models.User
	var userLoginInput UserLoginInput

	if err := c.ShouldBind(&userLoginInput); err != nil {
		return nil, jwt.ErrMissingLoginValues
	}

	if err := models.DB.Where("email = ?", userLoginInput.Email).Where("password = ?", userLoginInput.Password).First(&user).Error; err != nil {
		return nil, jwt.ErrFailedAuthentication
	}

	return &services.LoggedUser{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
		Name:  user.FirstName + " " + user.LastName,
	}, nil
}
