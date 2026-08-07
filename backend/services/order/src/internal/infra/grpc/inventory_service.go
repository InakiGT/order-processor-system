package grpc

import (
	"context"

	pb "github.com/InakiGT/processor/order-service/src/api/pb/product_stock/v1"
	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"github.com/InakiGT/processor/order-service/src/internal/domain/errors"
)

type InventoryServiceGRPC struct {
	client pb.ProductStockServiceClient
}

func NewInventoryService(client pb.ProductStockServiceClient) *InventoryServiceGRPC {
	return &InventoryServiceGRPC{client}
}

func (s *InventoryServiceGRPC) ReserveStock(ctx context.Context, productsIds []*entities.OrderItem) error {

	res, err := s.client.ReserveStock(ctx, toReserveStocks(productsIds))

	if res.Status == false {
		return errors.ErrInsuficientStock
	}

	if err != nil {
		return err
	}

	return nil
}
