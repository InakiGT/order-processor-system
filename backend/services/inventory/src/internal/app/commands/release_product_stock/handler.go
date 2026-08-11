package releaseproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type ReleaseProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewReleaseProductStock(repo repositories.ProductStockRepository) *ReleaseProductStockHandler {
	return &ReleaseProductStockHandler{repo}
}

func (h *ReleaseProductStockHandler) Handle(ctx context.Context, cmd ReleaseProductStockCommand) error {
	productsToRelease := make([]*entities.ProductStock, 0, len(cmd.Products))

	for _, item := range cmd.Products {
		product, err := h.repo.FindOneByID(ctx, item.ProductId)
		if err != nil {
			return err
		}

		if err := product.Release(item.Quantity); err != nil {
			return nil
		}

		productsToRelease = append(productsToRelease, product)
	}

	if err := h.repo.SaveMany(ctx, productsToRelease); err != nil {
		return err
	}

	return nil
}
