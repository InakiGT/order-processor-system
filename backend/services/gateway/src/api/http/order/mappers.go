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
		Msg: "Order created",
		Data: Order{
			OrderId:     int(res.Order.Id),
			Status:      res.Order.Status,
			TotalAmount: int(res.Order.TotalAmount),
		},
	}
}
