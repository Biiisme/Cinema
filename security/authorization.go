// security/authorization.go
package security

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your_secret_key")

// JWTAuthMiddleware kiểm tra token JWT hợp lệ và lấy thông tin người dùng
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "MEMBER") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ hoặc không được cung cấp"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "MEMBER")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token không hợp lệ"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Dữ liệu token không hợp lệ"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}

// AdminOnlyMiddleware cho phép truy cập chỉ dành cho người dùng admin
func AdminOnlyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Không được phép: Chỉ dành cho Admin"})
			c.Abort()
			return
		}
		c.Next()
	}
}
