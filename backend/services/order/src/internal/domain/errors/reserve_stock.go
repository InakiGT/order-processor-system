package errors

import "errors"

type ReserveStockErr error

var (
	ErrInsuficientStock ReserveStockErr = errors.New("Error while trying to reserver stock. Insuficient stock.")
)
