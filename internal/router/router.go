package router

import (
	"product_api/internal/handlers"
	"product_api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(productHandler *handlers.ProductHandler, authHandler *handlers.AuthHandler, cartHandler *handlers.CartHandler) *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        // Маршруты авторизации
        auth := api.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
        }

        // Маршруты для администраторов
        admin := api.Group("/admin", middleware.RoleMiddleware("admin"))
        {
            admin.POST("/products", productHandler.CreateProduct)
            admin.PUT("/products/:id", productHandler.UpdateProduct)
            admin.DELETE("/products/:id", productHandler.DeleteProduct)
        }

        cart := api.Group("/cart", middleware.AuthMiddleware())
        {
            cart.POST("", cartHandler.AddToCart)
            cart.GET("", cartHandler.GetCart)
            cart.PUT("/:id", cartHandler.UpdateCartItem)
            cart.DELETE("/:id", cartHandler.RemoveFromCart)
        }

        // Маршруты для всех пользователей
        api.GET("/products", productHandler.GetAllProducts)
        api.GET("/products/:id", productHandler.GetProductByID)
    }
    

    return r
}
