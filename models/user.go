package models

var UserIdkentityKey string = "id"

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
}
