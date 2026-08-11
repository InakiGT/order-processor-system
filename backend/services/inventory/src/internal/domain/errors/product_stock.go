package errors

import "errors"

var (
	ErrInvalidQuantity            = errors.New("Quantity must be greater than zero")
	ErrInsufficientStock          = errors.New("Insufficient available stock")
	ErrInsufficientStockToRelease = errors.New("Insufficient stock to release")
)
