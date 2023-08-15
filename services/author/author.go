package author

import (
	"bookstore/models"
	"time"
)

type CreateAuthorInput struct {
	FirstName string    `form:"first_name" json:"first_name" binding:"required"`
	LastName  string    `form:"last_name" json:"last_name" binding:"required"`
	BirthDay  time.Time `form:"birth_day" json:"birth_day" binding:"required"`
}

func Create(
	authorData CreateAuthorInput,
) (author models.Author) {
	author = models.Author{
		FirstName: authorData.FirstName,
		LastName:  authorData.LastName,
		BirthDay:  authorData.BirthDay,
	}
	models.DB.Create(&author)
	return author
}
