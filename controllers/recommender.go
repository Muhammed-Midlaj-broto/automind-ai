package controllers

import (
	"automind-ai/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecommendPage(c *gin.Context) {
	c.HTML(http.StatusOK, "recommender.html", nil)
}
func RecommenderAsk(c *gin.Context) {
	var req struct {
		Budget  string `json:"budget"`
		UseCase string `json:"useCase"`
		Terrain string `json:"terrain`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	prompt := "Suggest vehicles under" + req.Budget +
		"suitable for" + req.UseCase +
		"on" + req.Terrain + "terrain in India"
	response, err := helpers.SendAIRequest(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"answer": response})
}
