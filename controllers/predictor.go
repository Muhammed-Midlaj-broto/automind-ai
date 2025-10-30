package controllers

import (
	"automind-ai/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PredictorPage(c *gin.Context) {
	c.HTML(http.StatusOK, "predictor.html", nil)
}
func PredictorAsk(c *gin.Context) {
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	prompt := "Predict the performance of a car with these specs:\n"
	for k, v := range data {
		prompt += fmt.Sprintf("%s: %v\n", k, v)
	}

	response, err := helpers.SendAIRequest(prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response})
}
