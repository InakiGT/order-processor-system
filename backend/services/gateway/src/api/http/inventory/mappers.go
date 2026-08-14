package inventory

import productstockpb "github.com/InakiGT/processor/api-gateway/src/api/pb/product_stock/v1"

func toProductStockResponse(product *productstockpb.ProductStock) ProductStockResponse {
	return ProductStockResponse{
		ProductId:        product.Id,
		SKU:              product.Sku,
		Brand:            product.Brand,
		Model:            product.Model,
		AvailableStock:   product.AvailableStock,
		ReservedQuantity: product.ReservedQuantity,
	}
}

func toProductsStocksResponse(products []*productstockpb.ProductStock) ProductsStocksResponse {
	newProducts := make([]ProductStockResponse, 0, len(products))

	for _, product := range products {
		newProducts = append(newProducts, toProductStockResponse(product))
	}

	return ProductsStocksResponse{
		Products: newProducts,
	}
}
