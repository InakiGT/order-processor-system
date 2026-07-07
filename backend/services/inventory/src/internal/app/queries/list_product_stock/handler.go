package listproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type ListProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewListProductStock(repo repositories.ProductStockRepository) *ListProductStockHandler {
	return &ListProductStockHandler{repo}
}

func (h *ListProductStockHandler) Handle(ctx context.Context, query ListProductStockQuery) ([]*entities.ProductStock, error) {
	products, err := h.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}
