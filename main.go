package main

import (
	routes "crud-api/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
