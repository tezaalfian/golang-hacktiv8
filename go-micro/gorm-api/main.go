package main

import (
	"gorm-api/database"
	"gorm-api/routers"
)

func main() {
	database.StartDB()
	const PORT = ":8080"
	routers.StartServer().Run(PORT)
}
