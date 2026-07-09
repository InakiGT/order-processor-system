package reserveproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type ReserveProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewReserveProductStock(repo repositories.ProductStockRepository) *ReserveProductStockHandler {
	return &ReserveProductStockHandler{repo}
}

func (h *ReserveProductStockHandler) Handle(ctx context.Context, command ReserveProductStockCommand) error {
	product, err := h.repo.FindOneByID(ctx, command.ProductId)
	if err != nil {
		return err
	}

	if err := product.Reserve(command.Quantity); err != nil {
		return err
	}

	if _, err := h.repo.Save(ctx, product); err != nil {
		return err
	}

	return nil
}
