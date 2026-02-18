package postgres

import (
	"context"
	"os"
	"testing"

	"github.com/oTeeLeko/product-service/internal/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestProductRepository_Integration(t *testing.T) {
	dsn := os.Getenv("DB_SOURCE_TEST")
	if dsn == "" {
		t.Skip("Skipping integration test: DB_SOURCE_TEST not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	// Migrate
	db.AutoMigrate(&entity.Product{})

	repo := NewProductRepository(db)
	ctx := context.Background()

	// Test Create
	prod := &entity.Product{Name: "Integration Test", Price: 99}
	err = repo.Create(ctx, prod)
	if err != nil {
		t.Errorf("Create failed: %v", err)
	}

	// Test Update
	updates := map[string]interface{}{"name": "Updated Integration"}
	err = repo.Update(ctx, prod.ID, updates)
	if err != nil {
		t.Errorf("Update failed: %v", err)
	}
}
