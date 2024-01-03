package middleware

import (
	"github.com/gin-gonic/gin"
)

func Intercept(router *gin.Engine) {
	router.Use(crosMiddleware)
}

func crosMiddleware(c *gin.Context) {
	// Set the "Access-Control-Allow-Origin" header
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Authorization")
	// Optional: Set other CORS-related headers
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(200)
		return
	}
	c.Next()
}
