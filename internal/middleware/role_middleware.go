package middleware

import (
	"log"
	"net/http"
	"product_api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует или некорректный"})
			c.Abort()
			return
		}

		token = token[7:]
		log.Printf("Полученный токен: %s", token)

		claims, err := utils.ValidateJWT(token)
		if err != nil {
			log.Printf("Ошибка проверки токена: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		log.Printf("Роль пользователя: %s, Требуемая роль: %s", claims.Role, requiredRole)

		if claims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ запрещен"})
			c.Abort()
			return
		}

		c.Set("userId", claims.UserID)
		c.Next()
	}
}
