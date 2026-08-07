package reserveproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type ReserveProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewReserveProductStock(repo repositories.ProductStockRepository) *ReserveProductStockHandler {
	return &ReserveProductStockHandler{repo}
}

func (h *ReserveProductStockHandler) Handle(ctx context.Context, cmd ReserveProductStockCommand) error {
	productsToReserve := make([]*entities.ProductStock, 0, len(cmd.Products))

	for _, item := range cmd.Products {
		product, err := h.repo.FindOneByID(ctx, item.ProductId)

		if err != nil {
			return err
		}

		if err = product.Reserve(item.Quantity); err != nil {
			return err
		}

		productsToReserve = append(productsToReserve, product)
	}

	return h.repo.SaveMany(ctx, productsToReserve)
}
