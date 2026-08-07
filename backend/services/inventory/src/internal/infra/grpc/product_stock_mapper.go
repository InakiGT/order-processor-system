package grpc

import (
	productstockpb "github.com/InakiGT/processor/inventory-service/src/api/pb/product_stock/v1"
	reserveproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/reserve_product_stock"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
)

func toProductStock(entity *entities.ProductStock) *productstockpb.ProductStock {
	return &productstockpb.ProductStock{
		Id:               uint32(entity.ProductID),
		Sku:              string(entity.SKU),
		Brand:            entity.Brand,
		Model:            entity.Model,
		ReservedQuantity: int32(entity.ReservedQuantity),
		AvailableStock:   int32(entity.AvailableQuantity),
	}
}

func toProductStocks(entites []*entities.ProductStock) []*productstockpb.ProductStock {
	products := make([]*productstockpb.ProductStock, 0)

	for _, entity := range entites {
		products = append(products, toProductStock(entity))
	}

	return products
}

func toProductReserveStockCmd(products []*productstockpb.ReserveStockProduct) []reserveproductstock.ProductStockCommand {
	newProducts := make([]reserveproductstock.ProductStockCommand, 0)

	for _, product := range products {
		newProducts = append(newProducts, reserveproductstock.ProductStockCommand{
			ProductId: entities.ProductID(product.Id),
			Quantity:  int(product.Quantity),
		})
	}

	return newProducts
}
