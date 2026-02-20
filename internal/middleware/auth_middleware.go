package middleware

// import (
// 	"net/http"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"gin-app/internal/auth"
// )

// func AuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")

// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

// 		token, err := auth.ValidateToken(tokenString)
// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		claims := token.Claims.(map[string]interface{})
// 		c.Set("user_id", claims["user_id"])

// 		c.Next()
// 	}
// }
