package routers

import (
	"github.com/gin-gonic/gin"
	"sql-api/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	book := router.Group("/books")
	{
		book.POST("/", controllers.CreateBook)
		book.PUT("/:id", controllers.UpdateBook)
		book.GET("/:id", controllers.GetBook)
		book.DELETE("/:id", controllers.DeleteBook)
		book.GET("/", controllers.GetAllBooks)
	}

	return router
}
