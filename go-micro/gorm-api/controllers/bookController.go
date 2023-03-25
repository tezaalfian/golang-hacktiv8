package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm-api/database"
	"gorm-api/models"
	"net/http"
)

// CreateBook creates a new book
// @Summary Create a new book
// @Description Create a new book
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book object"
// @Success 201 {object} models.Book
// @Failure 400 {object} string
// @Router /books [post]
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

// UpdateBook updates a book
// @Summary Update a book
// @Description Update a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Param book body models.Book true "Book object"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Router /books/{id} [put]
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

// GetBook gets a book
// @Summary Get a book
// @Description Get a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} string
// @Router /books/{id} [get]
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

// GetAllBooks gets all books
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Book
// @Failure 400 {object} string
// @Router /books [get]
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

// DeleteBook deletes a book
// @Summary Delete a book
// @Description Delete a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path string true "Book ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /books/{id} [delete]
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
