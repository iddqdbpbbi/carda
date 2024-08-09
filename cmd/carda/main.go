package main

import (
	"carda/internal/carda/routes"
)

func main() {
	routes.SetupRoutes()
	routes.StartRouting()
}
