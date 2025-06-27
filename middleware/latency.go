package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LatencyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Last Request took %v", duration)
	}
}
