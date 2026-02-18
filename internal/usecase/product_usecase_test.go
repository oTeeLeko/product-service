package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/oTeeLeko/product-service/internal/adapter/handler/dto"
	"github.com/oTeeLeko/product-service/internal/domain/entity"
)

// Manual Mock Repository
type mockProductRepo struct {
	createFunc func(ctx context.Context, product *entity.Product) error
	updateFunc func(ctx context.Context, id uint, updates map[string]interface{}) error
}

func (m *mockProductRepo) Create(ctx context.Context, product *entity.Product) error {
	return m.createFunc(ctx, product)
}

func (m *mockProductRepo) Update(ctx context.Context, id uint, updates map[string]interface{}) error {
	return m.updateFunc(ctx, id, updates)
}

func TestProductUseCase_CreateProduct(t *testing.T) {
	mockRepo := &mockProductRepo{
		createFunc: func(ctx context.Context, product *entity.Product) error {
			product.ID = 1
			return nil
		},
	}

	useCase := NewProductUseCase(mockRepo)
	req := &dto.CreateProductRequest{
		Name:  "Test Product",
		Price: 100,
	}

	res, err := useCase.CreateProduct(context.Background(), req)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if res.ID != 1 {
		t.Errorf("Expected ID 1, got %d", res.ID)
	}
}

func TestProductUseCase_UpdateProduct_Error(t *testing.T) {
	mockRepo := &mockProductRepo{
		updateFunc: func(ctx context.Context, id uint, updates map[string]interface{}) error {
			return errors.New("update failed")
		},
	}

	useCase := NewProductUseCase(mockRepo)
	name := "Updated Name"
	req := &dto.UpdateProductRequest{
		Name: &name,
	}

	err := useCase.UpdateProduct(context.Background(), 1, req)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}
