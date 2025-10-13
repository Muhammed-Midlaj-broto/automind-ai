package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func RecommendPage(c *gin.Context) {
	c.HTML(http.StatusOK, "recommender.html", nil)
}

func AdvisorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "advisor.html", nil)
}
func PredictPage(c *gin.Context) {
	c.HTML(http.StatusOK, "predictor.html", nil)
}

func SummarizerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "summarizer.html", nil)
}
