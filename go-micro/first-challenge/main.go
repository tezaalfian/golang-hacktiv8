package main

import "first-challenge/routers"

func main() {
	const PORT = ":8080"

	routers.StartServer().Run(PORT)
}
