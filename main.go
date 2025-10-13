package main

import (
	"automind-ai/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Load all HTML templates
	r.LoadHTMLGlob("templates/*")

	// Serve static files (CSS, JS, images)
	r.Static("/static", "./static")

	// Setup routes
	routes.SetupRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
