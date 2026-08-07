package errors

import "errors"

type OrderError error

var (
	ErrEmptyInventory OrderError = errors.New("Inventory must have items")
)
