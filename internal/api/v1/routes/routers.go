package routes

import (
	"go_commerce/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, productHandler *handlers.ProductHandler) {
	v1 := r.Group("/api/v1")
	{
		// Product routes
		v1.POST("/product", productHandler.CreateProduct)
		v1.GET("/products/:id", productHandler.GetProduct)
		// v1.GET("/products", productHandler.GetAllProducts)
		// v1.PATCH("/products/:id", productHandler.UpdateProduct)
		// v1.DELETE("/products/:id", productHandler.DeleteProduct)
	}
}
