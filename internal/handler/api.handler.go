package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"hook-tail/internal/database"
	webhook "hook-tail/internal/model"
)

func CreateWebhookHandler(c *gin.Context) {
	var webhookData webhook.WebhookModel

	if err := c.ShouldBindJSON(&webhookData); err != nil {
		ResponseJSON(
			c,
			http.StatusBadRequest,
			"Invalid webhook data",
			nil,
		)
		return
	}

	if err := database.DB.Create(&webhookData).Error; err != nil {
		ResponseJSON(
			c,
			http.StatusConflict,
			"Webhook already processed",
			nil,
		)
		return
	}

	ResponseJSON(
		c,
		http.StatusCreated,
		"Webhook created successfully",
		webhookData,
	)
}
func GetWebhookHandler(c *gin.Context) {

}