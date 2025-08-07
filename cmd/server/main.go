package main

import (
	"go_commerce/internal/api/v1/handlers"
	"go_commerce/internal/api/v1/routes"
	"go_commerce/internal/domain/product"
	"go_commerce/internal/infrastructure/config"
	"go_commerce/internal/infrastructure/database"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to database
	db, err := database.NewConnection(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Initialize repositories
	productRepo := database.NewProductRepository(db)

	// Initialize services
	productService := product.NewService(productRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)

	// Setup Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, productHandler)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
