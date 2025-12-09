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

	// Run web serverport := os.Getenv("PORT")port := os.Getenv("PORT")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // for local testing
	}
	r.Run(":" + port)

}
