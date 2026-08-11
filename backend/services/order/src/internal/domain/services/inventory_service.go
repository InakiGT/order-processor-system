package services

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
)

type InventoryService interface {
	ReserveStock(ctx context.Context, productIds []*entities.OrderItem) error
	ReleaseStock(ctx context.Context, products []*entities.OrderItem) error
}
