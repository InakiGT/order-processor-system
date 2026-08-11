package releaseproductstock

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

type ReleaseProductStockCommand struct {
	Products []ProductStockCommand
}

type ProductStockCommand struct {
	ProductId entities.ProductID
	Quantity  int
}
