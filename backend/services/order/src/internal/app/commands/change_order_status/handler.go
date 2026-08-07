package changeorderstatus

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/repositories"
)

type ChangeOrderStatusHandler struct {
	repo repositories.OrderRepository
}

func NewChangeOrderStatus(repo repositories.OrderRepository) *ChangeOrderStatusHandler {
	return &ChangeOrderStatusHandler{repo}
}

func (h *ChangeOrderStatusHandler) Handle(ctx context.Context, cmd ChangeOrderStatusCommand) error {
	order, err := h.repo.FindOneById(ctx, cmd.OrderID)
	if err != nil {
		return err
	}

	order.Status = cmd.Status
	if _, err = h.repo.Save(ctx, order); err != nil {
		return err
	}

	return nil
}
