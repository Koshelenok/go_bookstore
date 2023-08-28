package book

import (
	"bookstore/models"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type CreateBookInput struct {
	Title    string `form:"title" binding:"required"`
	Author   string `form:"author" binding:"authorvalidator"`
	AuthorID uint   `form:"author_id" binding:"authorvalidator"`
}

type UpdateBookInput struct {
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

func GetById(id int) (book *models.Book, er error) {
	if err := models.DB.Where("id = ?", id).Preload("Author").First(&book).Error; err != nil {
		return nil, fmt.Errorf("record not found")
	}
	return book, nil
}

func Create(
	in CreateBookInput,
) (book models.Book) {
	book = models.Book{
		Title:    in.Title,
		AuthorID: in.AuthorID,
	}
	models.DB.Create(&book)
	return book
}

func Update(
	id int,
	in UpdateBookInput,
) (book *models.Book, err error) {
	book, err = GetById(id)
	if err != nil {
		return nil, err
	}
	models.DB.
		Model(&book).
		Updates(in)

	return book, nil
}
