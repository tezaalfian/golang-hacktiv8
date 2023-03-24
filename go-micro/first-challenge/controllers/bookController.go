package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookData = []Book{}

func CreateBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	book.ID = "BOOK" + strconv.Itoa(len(BookData)+1)
	BookData = append(BookData, book)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookData {
		if book.ID == id {
			BookData[i] = updatedBook
			BookData[i].ID = id
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %s not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s has been updated", id),
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var book Book

	for _, item := range BookData {
		if item.ID == id {
			condition = true
			book = item
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %s not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false

	for i, book := range BookData {
		if book.ID == id {
			BookData = append(BookData[:i], BookData[i+1:]...)
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data not found",
			"error_message": fmt.Sprintf("Book with id %s not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %s has been deleted", id),
	})
}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": BookData,
	})
}
