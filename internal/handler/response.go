package handler

import "github.com/gin-gonic/gin"

func ResponseJSON(
	c *gin.Context,
	status int,
	message string,
	data any,
) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}