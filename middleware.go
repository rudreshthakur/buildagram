package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		fmt.Printf("[BuildGram] %s %s | %v\n",
			c.Request.Method,
			c.Request.URL.Path,
			latency,
		)
	}
}