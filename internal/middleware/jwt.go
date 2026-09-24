package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"

	"visitor_management_backend/internal/helper"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// ========================================
		// AMBIL AUTHORIZATION HEADER
		// ========================================

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Format token salah",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(
			strings.TrimPrefix(authHeader, "Bearer "),
		)

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak ditemukan",
			})
			c.Abort()
			return
		}

		// ========================================
		// PARSE JWT
		// ========================================

		claims := &helper.Claims{}

		token, err := jwt.ParseWithClaims(
			tokenString,
			claims,
			func(token *jwt.Token) (interface{}, error) {

				// Pastikan algoritma HMAC.
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token tidak valid",
			})
			c.Abort()
			return
		}

		// ========================================
		// VALIDASI EMPLOYEE ID
		// ========================================

		if claims.ID <= 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Employee ID pada token tidak valid",
			})
			c.Abort()
			return
		}

		// ========================================
		// SIMPAN DATA JWT KE GIN CONTEXT
		// ========================================
		//
		// Bisa digunakan oleh handler:
		//
		// employeeID, exists := c.Get("employee_id")
		//

		c.Set("employee_id", claims.ID)
		c.Set("employee_email", claims.Email)
		c.Set("employee_role", claims.Role)

		c.Next()
	}
}