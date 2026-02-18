package dto

type CreateProductRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	SalePrice   *float64 `json:"sale_price"`
}

type ProductResponse struct {
	ID          uint     `json:"id"`
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	SalePrice   *float64 `json:"sale_price"`
}

type UpdateProductRequest struct {
	Name        string   `json:"name"`
	Description *string  `json:"description"`
	Price       float64  `json:"price"`
	SalePrice   *float64 `json:"sale_price"`
}
