package usecase

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/oTeeLeko/product-service/internal/adapter/handler/dto"
	"github.com/oTeeLeko/product-service/internal/domain/entity"
	"github.com/oTeeLeko/product-service/internal/domain/repository"
)

type ProductUseCase interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error)
	UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) error
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (u *productUseCase) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	var product entity.Product
	if err := copier.Copy(&product, req); err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, &product); err != nil {
		return nil, err
	}

	var res dto.ProductResponse
	if err := copier.Copy(&res, &product); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *productUseCase) UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) error {
	updates := make(map[string]interface{})

	// Handles "undefined" by checking for nil pointers
	if req.Name != nil {
		updates["name"] = *req.Name
	}
	if req.Description != nil {
		updates["description"] = req.Description
	} // description is already *string
	if req.Price != nil {
		updates["price"] = *req.Price
	}
	if req.SalePrice != nil {
		updates["sale_price"] = req.SalePrice
	} // sale_price is already *float64

	return u.repo.Update(ctx, id, updates)
}
