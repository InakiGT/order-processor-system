package order

import pb "github.com/InakiGT/processor/api-gateway/src/api/pb/order/v1"

func toCreateOrderGrpcItem(item OrderItem) *pb.CreateOrderItem {
	return &pb.CreateOrderItem{
		ProductId:       item.ProductId,
		Quantity:        item.Quantity,
		PriceAtPurchase: item.PriceAtPurchase,
	}
}

func toCreateOrderGrpcItems(items []OrderItem) []*pb.CreateOrderItem {
	newItems := make([]*pb.CreateOrderItem, 0, len(items))

	for _, item := range items {
		newItems = append(newItems, toCreateOrderGrpcItem(item))
	}

	return newItems
}

func toCreateOrderJSONResponse(res *pb.CreateOrderResponse) *CreateOrderResponse {
	return &CreateOrderResponse{
		Order: Order{
			OrderId:     int(res.Order.Id),
			Status:      res.Order.Status,
			TotalAmount: int(res.Order.TotalAmount),
		},
	}
}

func toItemResponse(item *pb.OrderItem) OrderItemResponse {
	return OrderItemResponse{
		OrderId:         item.OrderId,
		ProductId:       item.ProductId,
		Quantity:        item.Quantity,
		PriceAtPurchase: item.PriceAtPurchase,
	}
}

func toItemsResponse(items []*pb.OrderItem) []OrderItemResponse {
	newItems := make([]OrderItemResponse, 0, len(items))

	for _, item := range items {
		newItems = append(newItems, toItemResponse(item))
	}

	return newItems
}

func toListOrdersResponse(res *pb.ListOrdersResponse) *ListOrdersResponse {
	orders := make([]GetOrderResponse, 0, len(res.Orders))

	for _, order := range res.Orders {
		orders = append(orders, *toGetOrderJSONResponse(order))
	}

	return &ListOrdersResponse{
		Orders: orders,
	}
}

func toGetOrderJSONResponse(order *pb.Order) *GetOrderResponse {
	return &GetOrderResponse{
		Order: Order{
			OrderId:     int(order.Id),
			Status:      order.Status,
			TotalAmount: int(order.TotalAmount),
		},
		Items: toItemsResponse(order.Items),
	}
}
