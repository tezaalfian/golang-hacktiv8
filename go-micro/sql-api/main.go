package main

import (
	"sql-api/config"
	"sql-api/routers"
)

func main() {
	config.ConnectDB()
	const PORT = ":8080"
	routers.StartServer().Run(PORT)
}
