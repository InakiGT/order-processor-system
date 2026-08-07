package createorder

type CreateOrderCommand struct {
	Items []*CreateOrderItemCommand
}

type CreateOrderItemCommand struct {
	ProductID       int
	Quantity        int
	PriceAtPurchase int32
}
