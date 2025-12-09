package main

import (
	"automind-ai/helpers"
	"automind-ai/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("errorloading .env file")
	}
	// Initialize AI client first
	helpers.InitAI()

	// Test the connection (optional)
	response, err := helpers.SendAIRequest("Explain how a car engine works.")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("AI says:", response)
	}

	// Setup Gin
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	routes.SetupRoutes(r)

	// Run web server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // fallback for local
	}
	r.Run(":" + port)

}
