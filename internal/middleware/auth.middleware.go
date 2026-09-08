package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        token := c.GetHeader("Authorization")

        if token == "" {
            c.JSON(401, gin.H{
                "error": "unauthorized",
            })
            c.Abort()
            return
        }

        c.Next()
    }
}