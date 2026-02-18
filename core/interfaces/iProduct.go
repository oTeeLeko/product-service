package interfaces

import (
	"context"

	dto "github.com/oTeeLeko/product-service/core/models"
)

type IProductService interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error)
}
