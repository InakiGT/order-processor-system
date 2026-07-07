package repositories

import (
	"context"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
)

type ProductStockRepository interface {
	FindAll(ctx context.Context) ([]*entities.ProductStock, error)
	FindOneByID(ctx context.Context, id entities.ProductID) (*entities.ProductStock, error)
	Save(ctx context.Context, productStock *entities.ProductStock) (*entities.ProductStock, error)
	Update(ctx context.Context, productStock *entities.ProductStock) error
	Delete(ctx context.Context, id entities.ProductID) error
}
