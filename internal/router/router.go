package router

import (
	"product_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(productHandler *handlers.ProductHandler) *gin.Engine {
    r := gin.Default()

    api := r.Group("/api")
    {
        products := api.Group("/products")
        {
            products.GET("", productHandler.GetAllProducts)
            products.GET("/:id", productHandler.GetProductByID)
            products.POST("", productHandler.CreateProduct)
            products.PUT("/:id", productHandler.UpdateProduct)
            products.DELETE("/:id", productHandler.DeleteProduct)
        }
    }

    return r
}
