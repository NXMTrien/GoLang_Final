package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware kiểm tra xem người dùng có được xác thực hay không
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Kiểm tra token hoặc thông tin xác thực ở đây
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Nếu token hợp lệ, cho phép truy cập tiếp
		c.Next()
	}
}
