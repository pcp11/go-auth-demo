package main

import (
	"books/inits"
	"books/models"
)

func init() {
	inits.LoadEnv()
	inits.DBInit()
}

func main() {
	inits.DB.AutoMigrate(&models.Book{})
	inits.DB.AutoMigrate(&models.Person{})
	inits.DB.AutoMigrate(&models.User{})
}
