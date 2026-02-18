package postgres

import (
	"context"

	"github.com/oTeeLeko/product-service/internal/domain/entity"
	"github.com/oTeeLeko/product-service/internal/domain/repository"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).Create(product).Error
}

func (r *productRepository) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return r.db.WithContext(ctx).Model(&entity.Product{}).Where("id = ?", id).Updates(updates).Error
}
