package order

type CreateOrderResponse struct {
	Order Order `json:"order"`
}

type Order struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	TotalAmount int    `json:"total_amount"`
}

type GetOrderResponse struct {
	Order Order               `json:"order"`
	Items []OrderItemResponse `json:"items"`
}

type OrderItemResponse struct {
	OrderId         int32 `json:"order_id"`
	ProductId       int32 `json:"product_id"`
	Quantity        int32 `json:"quantity"`
	PriceAtPurchase int32 `json:"price_at_purchase"`
}

type ListOrdersResponse struct {
	Orders []GetOrderResponse `json:"orders"`
}
