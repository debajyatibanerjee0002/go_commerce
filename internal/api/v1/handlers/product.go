package handlers

import (
	"go_commerce/internal/domain/product"
	"go_commerce/internal/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	productService *product.Service
}

func NewProductHandler(productService *product.Service) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req product.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request", err)
		return
	}

	product, err := h.productService.CreateProduct(c.Request.Context(), &req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product", err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")

	id, err := uuid.Parse(idStr)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err)
		return
	}

	product, err := h.productService.GetProduct(c.Request.Context(), id)

	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Failed to get product", err)
	}

	if product == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Product not found", nil)
	}

	utils.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)

}
