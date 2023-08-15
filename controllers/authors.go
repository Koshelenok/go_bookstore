package controllers

import (
	"bookstore/models"
	authorService "bookstore/services/author"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UpdateAuthorUnput struct {
	FirstName string    `form:"first_name"`
	LastName  string    `form:"last_name"`
	BirthDay  time.Time `form:"birth_day"`
}

func CreateAuthor(c *gin.Context) {
	var input authorService.CreateAuthorInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	author := authorService.Create(input)

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

func FindAuthor(c *gin.Context) { // Get model if exist
	var author models.Author

	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func UpdateAuthor(c *gin.Context) {
	var author models.Author

	if err := models.DB.Where("id = ?", c.Param("id")).First(&author).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateAuthorUnput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&author).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": author})
}
