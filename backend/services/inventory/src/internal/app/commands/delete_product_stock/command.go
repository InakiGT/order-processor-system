package deleteproductstock

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

type DeleteProductStockCommand struct {
	ProductId entities.ProductID
}
