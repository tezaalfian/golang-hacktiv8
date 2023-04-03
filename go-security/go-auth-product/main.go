package main

import (
	"go-auth-product/database"
	"go-auth-product/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
