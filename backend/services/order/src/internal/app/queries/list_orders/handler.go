package listorders

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/order-service/src/internal/domain/repositories"
)

type ListOrdersHandler struct {
	repo repositories.OrderRepository
}

func NewListOrders(repo repositories.OrderRepository) *ListOrdersHandler {
	return &ListOrdersHandler{repo}
}

func (h *ListOrdersHandler) Handle(ctx context.Context, query ListOrdersQuery) ([]*entities.Order, error) {
	orders, err := h.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
