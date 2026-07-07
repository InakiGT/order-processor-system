package createproductstock

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/repositories"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/services"
)

type CreateProductStockHandler struct {
	repo repositories.ProductStockRepository
}

func NewCreateProductStock(repo repositories.ProductStockRepository) *CreateProductStockHandler {
	return &CreateProductStockHandler{repo}
}

func (h *CreateProductStockHandler) Handle(ctx context.Context, command CreateProductStockCommand) (entities.ProductID, error) {
	sku := services.GenerateSKU(command.Brand, command.Model)
	product := entities.NewProductStock(sku, command.Brand, command.Model, command.AvailableQuantity)

	savedP, err := h.repo.Save(ctx, product)
	if err != nil {
		return 0, err
	}

	return savedP.ProductID, nil
}
