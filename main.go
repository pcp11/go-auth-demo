package main

import (
	"books/controllers"
	"books/inits"
	"books/middlewares"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.AllowAllOrigins = false
	config.AllowOriginFunc = func(origin string) bool { return true }
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))
	r.GET("/books", middlewares.RequireAuth, controllers.GetBooks)
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
