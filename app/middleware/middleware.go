package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Intercept(router *gin.Engine) {
	router.Use(crosMiddleware)
	router.Use(jwtMiddleware)
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

func jwtMiddleware(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.Next()
		return
	}

	// Get the token after Bearer from the Authorization header
	parsedToken, err := jwt.Parse(token[7:], func(token *jwt.Token) (interface{}, error) {
		// Provide the secret key used for signing the token
		return []byte("9$(;:+q}3n@k:d7"), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"error": "Failed to parse JWT token"})
		c.Abort()
		return
	}

	// Check if the token is valid
	if parsedToken.Valid {
		// Extract the email from the token claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		email, ok := claims["email"].(string)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid email claim"})
			c.Abort()
			return
		}

		ctx := context.WithValue(c.Request.Context(), "email", &email)
		c.Request = c.Request.WithContext(ctx)

	} else {
		c.JSON(401, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Next()
}
