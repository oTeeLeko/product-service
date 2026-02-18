package store

import (
	"github.com/oTeeLeko/product-service/core/interfaces"
	"github.com/oTeeLeko/product-service/core/repository"
	"gorm.io/gorm"
)

type Store struct {
	ProductRepo interfaces.IProductService
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		ProductRepo: repository.NewProductRepository(db),
	}
}
