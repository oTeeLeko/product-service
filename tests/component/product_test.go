package component

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/oTeeLeko/product-service/internal/adapter/handler/dto"
	"github.com/oTeeLeko/product-service/internal/infrastructure/container"
	"github.com/oTeeLeko/product-service/internal/infrastructure/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestProductComponent_PostAndPatch(t *testing.T) {
	// Setup (Using a Mock DB or real Test DB)
	// For demonstration, we'll try to connect but skip if failing
	dsn := "postgresql://postgres:postgres@localhost:5432/product_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Skip("Skipping component test: Database not available")
	}

	c := container.NewContainer(db)
	r := router.SetupRouter(c)

	// 1. Test POST /product
	reqBody, _ := json.Marshal(dto.CreateProductRequest{
		Name:  "Component Test Product",
		Price: 200,
	})
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected 201, got %d", w.Code)
	}

	// 2. Test PATCH /product/1
	patchBody, _ := json.Marshal(map[string]interface{}{
		"name": "Patched by Component Test",
	})
	req, _ = http.NewRequest("PATCH", "/product/1", bytes.NewBuffer(patchBody))
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}
