package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/internal/adapter/handler/dto"
	"github.com/oTeeLeko/product-service/internal/usecase"
	"github.com/oTeeLeko/product-service/utils"
)

type ProductHandler struct {
	useCase usecase.ProductUseCase
}

func NewProductHandler(useCase usecase.ProductUseCase) *ProductHandler {
	return &ProductHandler{useCase: useCase}
}

// @Summary Create Product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product data"
// @Success 201 {object} utils.APIResponse{data=dto.ProductResponse} "Product created successfully"
// @Router /product [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	result, err := h.useCase.CreateProduct(ctx, &req)
	if err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	// Restore ORIGINAL response logic
	ctx.JSON(http.StatusCreated, utils.NewSuccessResponse(result))
}

// @Summary Update Product
// @Description Update an existing product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body dto.UpdateProductRequest true "Product data"
// @Success 200 {object} dto.SimpleResponse "Product updated successfully"
// @Router /product/{id} [patch]
func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		status := http.StatusBadRequest
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	var req dto.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	err = h.useCase.UpdateProduct(ctx, uint(id), &req)
	if err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	// PATCH still follows the 2-field spec as per global requirement
	ctx.JSON(http.StatusOK, dto.SimpleResponse{
		Successful: true,
		ErrorCode:  "",
	})
}
