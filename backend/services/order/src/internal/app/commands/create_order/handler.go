package createorder

import (
	"context"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/order-service/src/internal/domain/repositories"
	"github.com/InakiGT/processor/order-service/src/internal/domain/services"
)

type CreateOrderHandler struct {
	repo             repositories.OrderRepository
	inventoryService services.InventoryService
}

func NewCreateOrder(repo repositories.OrderRepository, service services.InventoryService) *CreateOrderHandler {
	return &CreateOrderHandler{repo, service}
}

func (h *CreateOrderHandler) Handle(ctx context.Context, cmd CreateOrderCommand) (*entities.Order, error) {
	items := make([]*entities.OrderItem, 0)

	for _, item := range cmd.Items {
		orderItem, err := entities.NewOrderItem(item.ProductID, item.Quantity, item.PriceAtPurchase)

		if err != nil {
			return nil, err
		}

		items = append(items, orderItem)
	}

	order, err := entities.NewOrder(items)
	if err != nil {
		return nil, err
	}

	if err := h.inventoryService.ReserveStock(ctx, items); err != nil {
		return nil, err
	}

	savedOrder, err := h.repo.Save(ctx, order)
	if err != nil {
		return nil, err
	}

	return savedOrder, nil
}
