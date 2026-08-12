package order

type CreateOrderRequest struct {
	items []OrderItem
}

type OrderItem struct {
	ProductId       int32 `json:"product_id"`
	Quantity        int32 `json:"quantity"`
	PriceAtPurchase int32 `json:"price_at_purchase"`
}
