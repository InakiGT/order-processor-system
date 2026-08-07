package gateways

import (
	orderpb "github.com/InakiGT/processor/order-service/src/api/pb/order/v1"
	stockpb "github.com/InakiGT/processor/order-service/src/api/pb/product_stock/v1"
	changeorderstatus "github.com/InakiGT/processor/order-service/src/internal/app/commands/change_order_status"
	createorder "github.com/InakiGT/processor/order-service/src/internal/app/commands/create_order"
	deleteorder "github.com/InakiGT/processor/order-service/src/internal/app/commands/delete_order"
	getorder "github.com/InakiGT/processor/order-service/src/internal/app/queries/get_order_by_id"
	listorders "github.com/InakiGT/processor/order-service/src/internal/app/queries/list_orders"
	infra "github.com/InakiGT/processor/order-service/src/internal/infra/grpc"
	x "github.com/InakiGT/processor/order-service/src/internal/infra/grpc"
	persistence "github.com/InakiGT/processor/order-service/src/internal/infra/persistence/gorm"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterOrderService(grpcServer *grpc.Server, db *gorm.DB, conn *grpc.ClientConn) {
	repo := persistence.NewOrderRepository(db)
	service := x.NewInventoryService(
		stockpb.NewProductStockServiceClient(conn),
	)

	get := getorder.NewGerOrderByID(repo)
	list := listorders.NewListOrders(repo)
	delete := deleteorder.NewDeleteOrder(repo)
	create := createorder.NewCreateOrder(repo, service)
	changeStatus := changeorderstatus.NewChangeOrderStatus(repo)

	handler := infra.NewOrderService(
		get,
		list,
		create,
		delete,
		changeStatus,
	)

	orderpb.RegisterOrderServiceServer(grpcServer, handler)
}
