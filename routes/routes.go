package routes

import (
	"automind-ai/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Home Page
	r.GET("/", controllers.HomePage)

	// Predictor Page
	r.GET("/predictor", controllers.PredictorPage)
	r.POST("/predictor/ask", controllers.PredictorAsk)

	// Recommender Page
	r.GET("/recommender", controllers.RecommendPage)
	r.POST("/recommender", controllers.RecommenderAsk)

	// Summarizer Page
	r.GET("/summarizer", controllers.SummarizerPage)
	r.POST("/summarizer/ask", controllers.SummarizerAsk)

	// Advisor Page
	r.GET("/advisor", controllers.AdvisorPage)
	r.POST("/advisor/ask", controllers.AdvisorAsk)
}
