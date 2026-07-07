package entities

import "time"

type ProductID uint
type SKU string

type ProductStock struct {
	ProductID         ProductID
	SKU               SKU
	Brand             string
	Model             string
	AvailableQuantity int
	ReservedQuantity  int
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

func NewProductStock(sku SKU, brand, model string, availableQuantity int) *ProductStock {
	return &ProductStock{
		SKU:               sku,
		Brand:             brand,
		Model:             model,
		AvailableQuantity: availableQuantity,
		ReservedQuantity:  0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
}
