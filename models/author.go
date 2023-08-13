package models

import "time"

type Author struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDay  time.Time `json:"birth_day"`
	Books     []Book    `gorm:"foreignKey:AuthorID"`
}
