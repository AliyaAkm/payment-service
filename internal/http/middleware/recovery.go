package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Recover() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		log.Printf("panic recovered: %v", recovered)
		c.AbortWithStatusJSON(500, gin.H{
			"error": "internal server error",
		})
	})
}
