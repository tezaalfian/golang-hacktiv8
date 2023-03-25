package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm-api/controllers"
	_ "gorm-api/docs"
)

// @title Book API
// @description This is a sample server for a book store.
// @version 1.0.0
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	book := router.Group("/books")
	{
		// Create
		book.POST("/", controllers.CreateBook)
		// Update
		book.PUT("/:id", controllers.UpdateBook)
		// Read
		book.GET("/:id", controllers.GetBook)
		// Delete
		book.DELETE("/:id", controllers.DeleteBook)
		// Read all
		book.GET("/", controllers.GetAllBooks)
	}
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
