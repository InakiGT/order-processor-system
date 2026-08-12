package order

type CreateOrderResponse struct {
	Msg  string `json:"msg"`
	Data Order  `json:"data"`
}

type Order struct {
	OrderId     int    `json:"order_id"`
	Status      string `json:"status"`
	TotalAmount int    `json:"total_amount"`
}
