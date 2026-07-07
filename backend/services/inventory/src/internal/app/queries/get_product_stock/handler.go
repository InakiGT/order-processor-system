package getproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
)

type GetProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewGetProductStock(repo repositories.ProductStockRepository) *GetProductStockHandler {
	return &GetProductStockHandler{repo}
}

func (h *GetProductStockHandler) Handle(ctx context.Context, query GetProductStockQuery) (*entities.ProductStock, error) {
	product, err := h.repo.FindOneByID(ctx, query.ProductId)
	if err != nil {
		return nil, err
	}

	return product, nil
}
