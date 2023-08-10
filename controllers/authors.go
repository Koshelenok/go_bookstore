package controllers

import (
	"bookstore/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateAuthorInput struct {
	FirstName string    `form:"first_name" binding:"required"`
	LastName  string    `form:"last_name" binding:"required"`
	BirthDay  time.Time `form:"birth_day" binding:"required"`
}

func CreateAuthor(c *gin.Context) {
	var input CreateAuthorInput

	if err := c.Bind(&input); err != nil {
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

func DeleteAuthor(c *gin.Context) {
	var author models.Author
	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&author)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
