package controllers

import (
	"bookstore/models"
	authorService "bookstore/services/author"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateBookInput struct {
	Title    string `form:"title" binding:"required"`
	Author   string `form:"author" binding:"authorvalidator"`
	AuthorID uint   `form:"author_id" binding:"authorvalidator"`
}

type UpdateBookUnput struct {
	Title  string `form:"title"`
	Author string `form:"author"`
}

var AuthorValidator validator.Func = func(fl validator.FieldLevel) bool {
	input := fl.Parent().Interface().(CreateBookInput)
	if input.AuthorID == 0 && !json.Valid([]byte(input.Author)) {
		return false
	}
	return true
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.AuthorID == 0 {
		var authorInput CreateAuthorInput
		json.Unmarshal([]byte(input.Author), &authorInput)
		author := authorService.Create(
			authorInput.FirstName,
			authorInput.LastName,
			authorInput.BirthDay,
		)
		input.AuthorID = author.ID
	}

	book := models.Book{Title: input.Title, AuthorID: input.AuthorID}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateBookUnput
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func FindBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func FindBook(c *gin.Context) { // Get model if exist
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}
