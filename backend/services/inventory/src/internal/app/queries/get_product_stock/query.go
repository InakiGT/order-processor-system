package getproductstock

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

type GetProductStockQuery struct {
	ProductId entities.ProductID
}
