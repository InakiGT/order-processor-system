package restockproductstock

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

type RestockProductStockCommand struct {
	ProductID entities.ProductID
	Quantity  int
}
