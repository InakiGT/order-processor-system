package repositories

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
)

type OrderRepository interface {
	FindAll(ctx context.Context) ([]*entities.Order, error)
	FindOneById(ctx context.Context, id entities.OrderID) (*entities.Order, error)
	Save(ctx context.Context, order *entities.Order) (*entities.Order, error)
	Delete(ctx context.Context, id entities.OrderID) error
}
