package tbl

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string   `gorm:"size:255;not null"`
	Description *string  `gorm:"size:512"`
	Price       float64  `gorm:"not null"`
	SalePrice   *float64 `gorm:"column:sale_price"`
}
