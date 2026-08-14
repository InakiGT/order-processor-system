package inventory

type ProductStockResponse struct {
	ProductId        uint32 `json:"product_id"`
	SKU              string `json:"sku"`
	Brand            string `json:"brand"`
	Model            string `json:"model"`
	AvailableStock   int32  `json:"available_stock"`
	ReservedQuantity int32  `json:"reserver_quantity"`
}

type ProductsStocksResponse struct {
	Products []ProductStockResponse `json:"products"`
}
