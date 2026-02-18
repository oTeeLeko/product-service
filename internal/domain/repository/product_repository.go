package repository

import (
	"context"

	"github.com/oTeeLeko/product-service/internal/domain/entity"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	Update(ctx context.Context, id uint, updates map[string]interface{}) error
}
