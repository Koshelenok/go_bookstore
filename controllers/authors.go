package controllers

import (
	"bookstore/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAuthorInput struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	BirthDay  time.Time `json:"birth_day" binding:"required"`
}

func CreateAuthor(c *gin.Context) {
	var input CreateAuthorInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		BirthDay:  input.BirthDay,
	}
	models.DB.Create(&author)

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func FindAuthors(c *gin.Context) {
	var authors []models.Author
	models.DB.Find(&authors)

	c.JSON(http.StatusOK, gin.H{"data": authors})
}
