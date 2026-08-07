package getorder

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/order-service/src/internal/domain/repositories"
)

type GetOrderByIDHandler struct {
	repo repositories.OrderRepository
}

func NewGerOrderByID(repo repositories.OrderRepository) *GetOrderByIDHandler {
	return &GetOrderByIDHandler{repo}
}

func (h *GetOrderByIDHandler) Handle(ctx context.Context, query GerOrderByIDQuery) (*entities.Order, error) {
	order, err := h.repo.FindOneById(ctx, query.OrderID)
	if err != nil {
		return nil, err
	}

	return order, nil
}
