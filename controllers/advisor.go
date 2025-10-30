package controllers

import (
	"automind-ai/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdvisorPage(c *gin.Context) {
	response, err := helpers.SendAIRequest("Give car maintenance tips for monsoon season")
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.HTML(http.StatusOK, "advisor.html", gin.H{
		"aiResponse": response,
	})
}
func AdvisorAsk(c *gin.Context) {
	var req struct {
		Prompt string `json:"prompt"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := helpers.SendAIRequest(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": response})
}
