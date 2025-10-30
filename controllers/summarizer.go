package controllers

import (
	"automind-ai/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SummarizerPage(c *gin.Context) {
	c.HTML(http.StatusOK, "summarizer.html", nil)
}
func SummarizerAsk(c *gin.Context) {
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		c.String(http.StatusBadRequest, "Invalid request")
		return
	}

	carModel := data["carModel"].(string)
	prompt := fmt.Sprintf("Summarize the key features, performance, and unique points of the car model: %s", carModel)

	response, err := helpers.SendAIRequest(prompt)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, response)
}
