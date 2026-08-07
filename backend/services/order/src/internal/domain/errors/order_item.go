package errors

import "errors"

type OrderItemError error

var (
	ErrInvalidQuantity OrderItemError = errors.New("Product quantity must be greater than 0")
	ErrInvalidPrice    OrderItemError = errors.New("Product price at purchase must be 0 or greater")
)
