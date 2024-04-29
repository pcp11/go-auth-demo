package main

import (
	"books/controllers"
	"books/inits"
	"books/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	r.Use(middlewares.CORSMiddleware())
	r.GET("/books", controllers.GetBooks)
	r.GET("/books/:id", controllers.GetBook)
	r.POST("/books", middlewares.RequireAuth, controllers.CreateBook)
	r.PUT("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)
	r.POST("/auth", middlewares.RequireAuth, controllers.Validate)

	err := inits.ImportBooks()
	if err != nil {
		log.Fatalf("Failed to import books: %v", err)
	}

	r.Run()
}
