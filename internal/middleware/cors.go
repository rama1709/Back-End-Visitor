package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware mengizinkan request lintas origin dari Frontend (misal Vite di
// http://localhost:3000). Origin yang diizinkan bisa diatur lewat env
// CORS_ALLOWED_ORIGINS (pisahkan dengan koma). Default sudah mencakup port dev
// yang umum dipakai TanStack/Vite.
func CORSMiddleware() gin.HandlerFunc {
	allowed := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowed == "" {
		allowed = "http://localhost:3000,http://127.0.0.1:3000,http://localhost:5173,http://127.0.0.1:5173"
	}
	allowedOrigins := strings.Split(allowed, ",")

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		for _, o := range allowedOrigins {
			if strings.TrimSpace(o) == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		c.Writer.Header().Set("Vary", "Origin")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
