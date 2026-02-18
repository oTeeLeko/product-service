package entity

import (
	"testing"
)

func TestProductEntity_Validation(t *testing.T) {
	desc := "Test Description"
	p := Product{
		Name:        "Test",
		Description: &desc,
		Price:       50.5,
	}

	if p.Name != "Test" {
		t.Errorf("Expected Test, got %s", p.Name)
	}
	if *p.Description != "Test Description" {
		t.Errorf("Expected description match")
	}
	if p.Price != 50.5 {
		t.Errorf("Expected 50.5, got %f", p.Price)
	}
}
