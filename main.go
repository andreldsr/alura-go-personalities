package main

import (
	"alura-go-personalities/database"
	"alura-go-personalities/routes"
)

func main() {
	database.Connect()
	routes.HandleRequests()
}
