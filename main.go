package main

import "sql-and-go/routers"

func main() {
	defer routers.CloseDB()
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
