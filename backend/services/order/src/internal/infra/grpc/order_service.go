package grpc

import (
	"context"

	orderpb "github.com/InakiGT/processor/order-service/src/api/pb/order/v1"
	changeorderstatus "github.com/InakiGT/processor/order-service/src/internal/app/commands/change_order_status"
	createorder "github.com/InakiGT/processor/order-service/src/internal/app/commands/create_order"
	deleteorder "github.com/InakiGT/processor/order-service/src/internal/app/commands/delete_order"
	getorder "github.com/InakiGT/processor/order-service/src/internal/app/queries/get_order_by_id"
	listorders "github.com/InakiGT/processor/order-service/src/internal/app/queries/list_orders"
	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	get          *getorder.GetOrderByIDHandler
	list         *listorders.ListOrdersHandler
	create       *createorder.CreateOrderHandler
	delete       *deleteorder.DeleteOrderHandler
	changeStatus *changeorderstatus.ChangeOrderStatusHandler
}

func NewOrderService(
	get *getorder.GetOrderByIDHandler,
	list *listorders.ListOrdersHandler,
	create *createorder.CreateOrderHandler,
	delete *deleteorder.DeleteOrderHandler,
	changeStatus *changeorderstatus.ChangeOrderStatusHandler,
) *OrderService {
	return &OrderService{
		get:          get,
		list:         list,
		create:       create,
		delete:       delete,
		changeStatus: changeStatus,
	}
}

func (s *OrderService) ListOrders(
	ctx context.Context,
	req *orderpb.ListOrdersRequest,
) (*orderpb.ListOrdersResponse, error) {
	query := listorders.ListOrdersQuery{}

	orders, err := s.list.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &orderpb.ListOrdersResponse{
		Orders: toOrders(orders),
	}, nil
}

func (s *OrderService) GetOrder(
	ctx context.Context,
	req *orderpb.GetOrderRequest,
) (*orderpb.GetOrderResponse, error) {
	query := getorder.GerOrderByIDQuery{
		OrderID: entities.OrderID(req.Id),
	}

	order, err := s.get.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &orderpb.GetOrderResponse{
		Order: toOrder(order),
	}, nil
}

func (s *OrderService) CreateOrder(
	ctx context.Context,
	req *orderpb.CreateOrderRequest,
) (*orderpb.CreateOrderResponse, error) {
	cmd := createorder.CreateOrderCommand{
		Items: toCmdOrderItems(req.Items),
	}

	order, err := s.create.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &orderpb.CreateOrderResponse{
		Order: toOrder(order),
	}, nil
}

func (s *OrderService) DeleteOrder(
	ctx context.Context,
	req *orderpb.DeleteOrderRequest,
) (*emptypb.Empty, error) {
	cmd := deleteorder.DeleteOrderCommand{
		Id: entities.OrderID(req.Id),
	}

	err := s.delete.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *OrderService) ChangeOrderStatus(
	ctx context.Context,
	req *orderpb.ChangeOrderStatusRequest,
) (*emptypb.Empty, error) {
	cmd := changeorderstatus.ChangeOrderStatusCommand{
		OrderID: entities.OrderID(req.Id),
		Status:  entities.OrderStatus(req.Status),
	}

	if err := s.changeStatus.Handle(ctx, cmd); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
