package repository

import (
	"context"

	"github.com/jinzhu/copier"
	dto "github.com/oTeeLeko/product-service/core/models"
	tbl "github.com/oTeeLeko/product-service/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*dto.ProductResponse, error) {
	var entity tbl.Product
	if err := copier.Copy(&entity, req); err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Create(&entity).Error; err != nil {
		return nil, err
	}
	var response dto.ProductResponse
	if err := copier.Copy(&response, &entity); err != nil {
		return nil, err
	}
	return &response, nil
}

func (r *ProductRepository) UpdateProduct(ctx context.Context, id uint, req *dto.UpdateProductRequest) (*dto.ProductResponse, error) {
	var entity tbl.Product
	if err := r.db.WithContext(ctx).First(&entity, id).Error; err != nil {
		return nil, err
	}
	if err := copier.CopyWithOption(&entity, req, copier.Option{IgnoreEmpty: true}); err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Save(&entity).Error; err != nil {
		return nil, err
	}

	var response dto.ProductResponse
	if err := copier.Copy(&response, &entity); err != nil {
		return nil, err
	}
	return &response, nil
}
