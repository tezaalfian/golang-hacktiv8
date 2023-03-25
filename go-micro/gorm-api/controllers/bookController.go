package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-api/database"
	"gorm-api/models"
	"net/http"
)

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := database.GetDB().Create(&book).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := database.GetDB().Model(&models.Book{}).Where("id = ?", id).Updates(updatedBook).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Updated book %s successfully", updatedBook.NameBook),
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	err := database.GetDB().First(&book, id).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetAllBooks(ctx *gin.Context) {
	var books []models.Book
	err := database.GetDB().Find(&books).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	err := database.GetDB().Where("id = ?", id).Delete(&book).Error
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s deleted successfully", id),
	})
}
