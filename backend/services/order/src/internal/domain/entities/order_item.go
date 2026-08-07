package entities

import "github.com/InakiGT/processor/order-service/src/internal/domain/errors"

type OrderItemID int

type OrderItem struct {
	OrderItemID     OrderItemID
	OrderID         OrderID
	ProductID       int
	Quantity        int
	PriceAtPurchase int32
}

func NewOrderItem(productId, quantity int, priceAtPurchase int32) (*OrderItem, error) {
	if quantity <= 0 {
		return nil, errors.ErrInvalidQuantity
	}
	if priceAtPurchase < 0 {
		return nil, errors.ErrInvalidPrice
	}

	return &OrderItem{
		Quantity:        quantity,
		ProductID:       productId,
		PriceAtPurchase: priceAtPurchase,
	}, nil
}
