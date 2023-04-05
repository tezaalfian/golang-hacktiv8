package main

import (
	"go-auth-product/database"
	"go-auth-product/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
