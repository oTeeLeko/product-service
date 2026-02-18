package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/product-service/core/interfaces"
	dto "github.com/oTeeLeko/product-service/core/models"
	"github.com/oTeeLeko/product-service/utils"
)

type ProductController struct {
	productRepo interfaces.IProductService
}

func NewProductController(productRepo interfaces.IProductService) *ProductController {
	return &ProductController{productRepo: productRepo}
}

// @Summary Create Product
// @Description Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param request body dto.CreateProductRequest true "Product data"
// @Success 201 {object} utils.APIResponse{data=dto.ProductResponse} "Product created successfully"
// @Failure 400 {object} utils.APIResponse{data=nil} "Invalid JSON payload"
// @Failure 500 {object} utils.APIResponse{data=nil} "Internal server error"
// @Router /product [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req *dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	result, err := c.productRepo.CreateProduct(ctx, req)
	if err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	ctx.JSON(http.StatusCreated, utils.NewSuccessResponse(result))
}

// @Summary Update Product
// @Description Update an existing product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body dto.UpdateProductRequest true "Product data"
// @Success 200 {object} utils.APIResponse{data=dto.ProductResponse} "Product updated successfully"
// @Failure 400 {object} utils.APIResponse{data=nil} "Invalid URI parameter or JSON payload"
// @Failure 500 {object} utils.APIResponse{data=nil} "Internal server error"
// @Router /product/{id} [patch]
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	var uri dto.BaseURL
	var req *dto.UpdateProductRequest

	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	result, err := c.productRepo.UpdateProduct(ctx, uri.ID, req)
	if err != nil {
		ctx.Error(err)
		status := utils.ParseError(err)
		ctx.JSON(status, utils.NewErrorResponse(status))
		return
	}

	ctx.JSON(http.StatusOK, utils.NewSuccessResponse(result))
}
