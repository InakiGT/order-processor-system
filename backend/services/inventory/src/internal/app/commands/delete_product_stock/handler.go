package deleteproductstock

import (
	"context"
	"fmt"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type DeleteProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewDeleteProductStock(repo repositories.ProductStockRepository) *DeleteProductStockHandler {
	return &DeleteProductStockHandler{repo}
}

func (h *DeleteProductStockHandler) Handle(ctx context.Context, command DeleteProductStockCommand) error {
	product, err := h.repo.FindOneByID(ctx, command.ProductId)
	if err != nil {
		return err
	}

	if product == nil {
		return fmt.Errorf("Product Stock ID not found")
	}

	if err = h.repo.Delete(ctx, product.ProductID); err != nil {
		return err
	}

	return nil
}
