package models

type Book struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Title    string `json:"title"`
	AuthorID uint   `json:"author_id"`
	Author   Author `json:"author" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:AuthorID"`
}
