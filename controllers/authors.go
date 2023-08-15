package controllers

import (
	"bookstore/models"
	authorService "bookstore/services/author"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := authorService.GetById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
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
	var input authorService.UpdateAuthorInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, err := getIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author, err := authorService.Update(
		id,
		input,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": author})
}

func getIDParam(c *gin.Context) (id int, err error) {
	rawID := c.Param("id")
	parsedID, err := strconv.Atoi(rawID)
	if err != nil {
		return 0, fmt.Errorf("incorrect ID parameter")
	}
	return parsedID, nil
}
