package main

import (
	"bookstore/controllers"
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/author", controllers.FindAuthors)
	r.POST("/author", controllers.CreateAuthor)
	r.DELETE("/author/:id", controllers.DeleteAuthor)

	r.Run()
}
