package reserveproductstock

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

type ReserveProductStockCommand struct {
	ProductId entities.ProductID
	Quantity  int
}
