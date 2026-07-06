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

func (h *CreateProductStockHandler) Handle(ctx context.Context, command CreateProductStockCommand) {
	sku := services.GenerateSKU(command.Brand, command.Model)
	product := entities.NewProductStock(sku, command.Brand, command.Model, command.AvailableQuantity)

	h.repo.Save(ctx, product)
}
