package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		// simple console log, replace with zap if needed
		_ = status
		//fmt.Printf("[%d] %s %s %v\n", status, method, path, latency)
		_ = latency
		_ = method
		_ = path
	}
}
