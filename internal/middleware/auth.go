package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"

	"visitor_management_backend/internal/helper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ==================================================
		// AUTHORIZATION HEADER
		// ==================================================

		authHeader := strings.TrimSpace(
			c.GetHeader("Authorization"),
		)

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header tidak ditemukan",
			})
			c.Abort()
			return
		}

		// ==================================================
		// BEARER TOKEN
		// ==================================================

		parts := strings.Fields(authHeader)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Format token salah",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// ==================================================
		// JWT SECRET
		// ==================================================

		jwtSecret := strings.TrimSpace(
			os.Getenv("JWT_SECRET"),
		)

		if jwtSecret == "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "JWT_SECRET belum dikonfigurasi",
			})
			c.Abort()
			return
		}

		// ==================================================
		// PARSE JWT
		// ==================================================

		var claims helper.Claims

		token, err := jwt.ParseWithClaims(
			tokenString,
			&claims,
			func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(jwtSecret), nil
			},
		)

		// ==================================================
		// VALIDATE TOKEN
		// ==================================================

		if err != nil ||
			token == nil ||
			!token.Valid {

			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		// ==================================================
		// VALIDATE USER ID
		// ==================================================

		if claims.ID <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Informasi user pada token tidak valid",
			})
			c.Abort()
			return
		}

		// ==================================================
		// SET USER DATA
		// ==================================================

		c.Set("user_id", claims.ID)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)

		c.Next()
	}
}
