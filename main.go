package main

import (
	_ "grip.app.api/docs"
	"grip.app.api/internal/middleware"
	"log"
)

// @title Grip API
// @version 1.0
// @description This is a Grip service API.
// @host localhost:8080
// @BasePath /
func main() {
	if err := middleware.Run(); err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}
}
