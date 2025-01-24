package middleware

import (
	"net/http"
	"product_api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		token := c.GetHeader("Authorization")
		if token == "" || len(token) < 7 || token[:7] != "Bearer " {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует или некорректный"})
			c.Abort()
			return
		}

		// Убираем "Bearer " из токена
		token = token[7:]

		// Проверяем токен
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		// Устанавливаем userId в контекст
		c.Set("userId", claims.UserID)
		c.Next()
	}
}
