package grpc

import (
	orderpb "github.com/InakiGT/processor/order-service/src/api/pb/order/v1"
	productstockpb "github.com/InakiGT/processor/order-service/src/api/pb/product_stock/v1"
	createorder "github.com/InakiGT/processor/order-service/src/internal/app/commands/create_order"
	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
)

func toOrderItem(item *entities.OrderItem) *orderpb.OrderItem {
	return &orderpb.OrderItem{
		OrderId:         int32(item.OrderID),
		Quantity:        int32(item.Quantity),
		PriceAtPurchase: item.PriceAtPurchase,
		ProductId:       int32(item.ProductID),
	}
}

func toOrderItems(items []*entities.OrderItem) []*orderpb.OrderItem {
	newItems := make([]*orderpb.OrderItem, 0)

	for _, item := range items {
		newItems = append(newItems, toOrderItem(item))
	}

	return newItems
}

func toOrder(order *entities.Order) *orderpb.Order {
	return &orderpb.Order{
		Id:          int32(order.OrderId),
		TotalAmount: order.TotalAmount,
		Status:      string(order.Status),
		Items:       toOrderItems(order.Items),
	}
}

func toOrders(orders []*entities.Order) []*orderpb.Order {
	newOrders := make([]*orderpb.Order, 0)

	for _, order := range orders {
		newOrders = append(newOrders, toOrder(order))
	}

	return newOrders
}

func toCmdOrderItems(items []*orderpb.OrderItem) []*createorder.CreateOrderItemCommand {
	orderItems := make([]*createorder.CreateOrderItemCommand, 0)

	for _, item := range items {
		orderItems = append(orderItems, &createorder.CreateOrderItemCommand{
			Quantity:        int(item.Quantity),
			ProductID:       int(item.ProductId),
			PriceAtPurchase: item.PriceAtPurchase,
		})
	}

	return orderItems
}

func toReserveStocks(items []*entities.OrderItem) *productstockpb.ReserveStockRequest {
	products := make([]*productstockpb.ReserveStockProduct, 0)

	for _, item := range items {
		products = append(products, toReserveStock(item))
	}

	return &productstockpb.ReserveStockRequest{
		Products: products,
	}
}

func toReserveStock(item *entities.OrderItem) *productstockpb.ReserveStockProduct {
	return &productstockpb.ReserveStockProduct{
		Id:       uint32(item.ProductID),
		Quantity: int32(item.Quantity),
	}
}
