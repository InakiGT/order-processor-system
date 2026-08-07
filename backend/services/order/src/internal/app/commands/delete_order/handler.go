package deleteorder

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/repositories"
)

type DeleteOrderHandler struct {
	repo repositories.OrderRepository
}

func NewDeleteOrder(repo repositories.OrderRepository) *DeleteOrderHandler {
	return &DeleteOrderHandler{repo}
}

func (h *DeleteOrderHandler) Handle(ctx context.Context, cmd DeleteOrderCommand) error {
	if err := h.repo.Delete(ctx, cmd.Id); err != nil {
		return err
	}

	return nil
}
