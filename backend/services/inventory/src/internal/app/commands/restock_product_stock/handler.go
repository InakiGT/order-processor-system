package restockproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type RestockProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewRestockProductStock(repo repositories.ProductStockRepository) *RestockProductStockHandler {
	return &RestockProductStockHandler{repo}
}

func (h *RestockProductStockHandler) Handle(ctx context.Context, cmd RestockProductStockCommand) error {
	product, err := h.repo.FindOneByID(ctx, cmd.ProductID)
	if err != nil {
		return err
	}

	if err = product.Restock(cmd.Quantity); err != nil {
		return err
	}

	if _, err = h.repo.Save(ctx, product); err != nil {
		return err
	}

	return nil
}
