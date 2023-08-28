package author

import (
	"bookstore/models"
	"fmt"
	"time"
)

type CreateAuthorInput struct {
	FirstName string    `form:"first_name" json:"first_name" binding:"required"`
	LastName  string    `form:"last_name" json:"last_name" binding:"required"`
	BirthDay  time.Time `form:"birth_day" json:"birth_day" binding:"required"`
}

type UpdateAuthorInput struct {
	FirstName string    `form:"first_name"`
	LastName  string    `form:"last_name"`
	BirthDay  time.Time `form:"birth_day"`
}

func GetById(id int) (author *models.Author, err error) {
	if err := models.DB.Where("id = ?", id).Preload("Books").First(&author).Error; err != nil {
		return nil, fmt.Errorf("record not found")
	}
	return author, nil
}

func Create(
	in CreateAuthorInput,
) (author models.Author) {
	author = models.Author{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		BirthDay:  in.BirthDay,
	}
	models.DB.Create(&author)
	return author
}

func Update(
	id int,
	in UpdateAuthorInput,
) (author *models.Author, err error) {
	author, err = GetById(id)
	if err != nil {
		return nil, err
	}
	models.DB.
		Model(&author).
		Updates(in)
	return author, nil
}
