package main

import (
	"bookstore/controllers"
	"bookstore/models"
	"bookstore/services"
	bookService "bookstore/services/book"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	authMiddleware, _ := services.SetupAuth(
		controllers.LoginHandler,
		controllers.AuthorizatorHandler,
	)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("authorvalidator", bookService.AuthorValidator)
	}

	r.POST("/api/login", authMiddleware.LoginHandler)

	auth := r.Group("/api")

	// the jwt middleware
	// Refresh time can be longer than token timeout
	auth.GET("/api/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", controllers.HelloHandler)

		auth.GET("/books", controllers.FindBooks)
		auth.POST("/books", controllers.CreateBook)
		auth.GET("/books/:id", controllers.FindBook)
		auth.PATCH("/books/:id", controllers.UpdateBook)
		auth.DELETE("/books/:id", controllers.DeleteBook)

		auth.GET("/author", controllers.FindAuthors)
		auth.POST("/author", controllers.CreateAuthor)
		auth.GET("/author/:id", controllers.FindAuthor)
		auth.PATCH("/author/:id", controllers.UpdateAuthor)
		auth.DELETE("/author/:id", controllers.DeleteAuthor)
	}

	r.Run()
}
