package order

type CreateOrderRequest struct {
	Items []OrderItem `json:"items" binding:"required,min=1"`
}

type OrderItem struct {
	ProductId       int32 `json:"product_id" binding:"required"`
	Quantity        int32 `json:"quantity" binding:"required"`
	PriceAtPurchase int32 `json:"price_at_purchase" biding:"required"`
}

type ChangeOrderStatusRequest struct {
	OrderId int32  `json:"order_id" binding:"required"`
	Status  string `json:"status" binding:"required"`
}
