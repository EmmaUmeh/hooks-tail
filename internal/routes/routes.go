package routes


import (
    "github.com/gin-gonic/gin"
    "hook-tail/internal/handler"
)

func SetupRoutes(router *gin.Engine) {

	v1:= router.Group("/api/v1")

	v1.POST("/webhooks", handler.CreateWebhookHandler)
	v1.GET("/webhooks/:id", handler.GetWebhookHandler)
}