package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		start := time.Now()
		c.Next() // Process request
		duration := time.Since(start)

		log.Printf("[%s] %s %s %d %v", c.ClientIP(), c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	})
}
