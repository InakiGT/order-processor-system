package entities

import (
	"time"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/errors"
)

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

func (p *ProductStock) Reserve(quantity int) error {
	if quantity <= 0 {
		return errors.ErrInvalidQuantity
	}

	if quantity > p.AvailableQuantity {
		return errors.ErrInsufficientStock
	}

	p.AvailableQuantity -= quantity
	p.ReservedQuantity += quantity
	p.UpdatedAt = time.Now()

	return nil
}
