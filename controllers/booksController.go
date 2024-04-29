package controllers

import (
	"books/inits"
	"books/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	var books []models.Book

	result := inits.DB.Preload("Authors").Find(&books)
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": books})
}

func GetBook(ctx *gin.Context) {
	var book models.Book

	result := inits.DB.Preload("Authors").First(&book, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": book})
}

func CreateBook(ctx *gin.Context) {
	var body struct {
		Title  string `json:"title"`
		UserID uint   `json:"user_id"`
	}
	ctx.BindJSON(&body)

	user, exists := ctx.Get("user")

	if !exists {
		ctx.JSON(500, gin.H{"error": "user not found"})
		return
	}
	body.UserID = user.(models.User).ID
	book := models.Book{Title: body.Title}
	result := inits.DB.Create(&book)

	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	ctx.JSON(200, gin.H{"data": book})
}

func UpdateBook(ctx *gin.Context) {
	var body struct {
		Title string
	}
	ctx.BindJSON(&body)

	var book models.Book

	result := inits.DB.First(&book, ctx.Param("id"))
	if result.Error != nil {
		ctx.JSON(500, gin.H{"error": result.Error})
		return
	}
	inits.DB.Model(&book).Updates(models.Book{Title: body.Title})
	ctx.JSON(200, gin.H{"data": book})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	inits.DB.Delete(&models.Book{}, id)
	ctx.JSON(200, gin.H{"data": "post has been deleted successfully"})
}
