package author

import (
	"bookstore/models"
	"time"
)

func Create(
	firstName string,
	lastName string,
	birthDay time.Time,
) (author models.Author) {
	author = models.Author{
		FirstName: firstName,
		LastName:  lastName,
		BirthDay:  birthDay,
	}
	models.DB.Create(&author)
	return author
}
