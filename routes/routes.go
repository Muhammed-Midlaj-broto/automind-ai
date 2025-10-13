package routes

import (
	"automind-ai/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Home Page
	r.GET("/", controllers.HomePage)

	// Predictor Page
	r.GET("/predictor", controllers.PredictPage)

	// Recommender Page
	r.GET("/recommender", controllers.RecommendPage)

	// Summarizer Page
	r.GET("/summarizer", controllers.SummarizerPage)

	// Advisor Page
	r.GET("/advisor", controllers.AdvisorPage)
}
