package repositories

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
)

type ProductStockRepository interface {
	FindAll(ctx context.Context)
	FindOneByID(ctx context.Context, id entities.ProductID)
	Save(ctx context.Context, productStock *entities.ProductStock)
	Update(ctx context.Context, productStock *entities.ProductStock)
	Delete(ctx context.Context, id entities.ProductID)
}
