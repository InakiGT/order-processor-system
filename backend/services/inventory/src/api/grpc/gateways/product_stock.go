package gateways

import (
	productstockpb "github.com/InakiGT/processor/inventory-service/src/api/pb/product_stock/v1"
	create "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/create_product_stock"
	delete "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/delete_product_stock"
	get "github.com/InakiGT/processor/inventory-service/src/internal/app/queries/get_product_stock"
	list "github.com/InakiGT/processor/inventory-service/src/internal/app/queries/list_product_stock"
	infra "github.com/InakiGT/processor/inventory-service/src/internal/infra/grpc"
	persistence "github.com/InakiGT/processor/inventory-service/src/internal/infra/persistence/gorm"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func RegisterProductStockService(grpcServer *grpc.Server, db *gorm.DB) {
	repo := persistence.NewProductStockRepository(db)

	get := get.NewGetProductStock(repo)
	list := list.NewListProductStock(repo)
	create := create.NewCreateProductStock(repo)
	delete := delete.NewDeleteProductStock(repo)

	handler := infra.NewProductStockService(
		get,
		list,
		create,
		delete,
	)

	productstockpb.RegisterProductStockServiceServer(grpcServer, handler)
}
