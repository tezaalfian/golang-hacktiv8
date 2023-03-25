package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sql-api/config"
	"sql-api/models"
)

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	sqlStatement := `INSERT INTO books (title, author, description) VALUES ($1, $2, $3) RETURNING *`
	err := config.GetDB().QueryRow(sqlStatement, book.Title, book.Author, book.Description).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
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

	sqlStatement := `UPDATE books SET title = $2, author = $3, description = $4 WHERE id = $1 RETURNING *`
	res, err := config.GetDB().Exec(sqlStatement, id, updatedBook.Title, updatedBook.Author, updatedBook.Description)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Updated %d book", rowsAffected),
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var book models.Book
	sqlStatement := `SELECT * FROM books WHERE id = $1`
	err := config.GetDB().QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"book": book,
	})
}

func GetAllBooks(ctx *gin.Context) {
	var books = []models.Book{}

	sqlStatement := `SELECT * FROM books`
	rows, err := config.GetDB().Query(sqlStatement)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		books = append(books, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	sqlStatement := `DELETE FROM books WHERE id = $1`
	res, err := config.GetDB().Exec(sqlStatement, id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Deleted %d book", rowsAffected),
	})
}
