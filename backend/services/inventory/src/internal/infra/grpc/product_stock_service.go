package grpc

import (
	"context"

	productstockpb "github.com/InakiGT/processor/inventory-service/src/api/pb/product_stock/v1"
	createproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/create_product_stock"
	deleteproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/delete_product_stock"
	releaseproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/release_product_stock"
	reserveproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/reserve_product_stock"
	restockproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/commands/restock_product_stock"
	getproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/queries/get_product_stock"
	listproductstock "github.com/InakiGT/processor/inventory-service/src/internal/app/queries/list_product_stock"
	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ProductStockService struct {
	productstockpb.UnimplementedProductStockServiceServer

	getHandler     *getproductstock.GetProductStockHandler
	listHandler    *listproductstock.ListProductStockHandler
	createHandler  *createproductstock.CreateProductStockHandler
	deleteHandler  *deleteproductstock.DeleteProductStockHandler
	reserveHandler *reserveproductstock.ReserveProductStockHandler
	restockHandler *restockproductstock.RestockProductStockHandler
	releaseHandler *releaseproductstock.ReleaseProductStockHandler
}

func NewProductStockService(
	get *getproductstock.GetProductStockHandler,
	list *listproductstock.ListProductStockHandler,
	create *createproductstock.CreateProductStockHandler,
	delete *deleteproductstock.DeleteProductStockHandler,
	reserve *reserveproductstock.ReserveProductStockHandler,
	restock *restockproductstock.RestockProductStockHandler,
	release *releaseproductstock.ReleaseProductStockHandler,
) *ProductStockService {
	return &ProductStockService{
		getHandler:     get,
		listHandler:    list,
		createHandler:  create,
		deleteHandler:  delete,
		reserveHandler: reserve,
		restockHandler: restock,
		releaseHandler: release,
	}
}

func (s *ProductStockService) CreateProductStock(
	ctx context.Context,
	req *productstockpb.CreateProductStockRequest) (*productstockpb.CreateProductStockResponse, error) {
	cmd := createproductstock.CreateProductStockCommand{
		Brand:             req.Brand,
		Model:             req.Model,
		AvailableQuantity: int(req.AvailableStock),
	}

	id, err := s.createHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return &productstockpb.CreateProductStockResponse{
		Id: uint32(id),
	}, nil
}

func (s *ProductStockService) DeleteProductStock(
	ctx context.Context,
	req *productstockpb.DeleteProductStockRequest) (*emptypb.Empty, error) {
	cmd := deleteproductstock.DeleteProductStockCommand{
		ProductId: entities.ProductID(req.Id),
	}

	if err := s.deleteHandler.Handle(ctx, cmd); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *ProductStockService) ListProductStocks(
	ctx context.Context,
	req *productstockpb.ListProductStocksRequest,
) (*productstockpb.ListProductStocksResponse, error) {
	query := listproductstock.ListProductStockQuery{}

	products, err := s.listHandler.Handle(ctx, query)

	if err != nil {
		return nil, err
	}

	return &productstockpb.ListProductStocksResponse{
		Products: toProductStocks(products),
	}, nil
}

func (s *ProductStockService) GetProductStock(
	ctx context.Context,
	req *productstockpb.GetProductStockRequest,
) (*productstockpb.GetProductStockResponse, error) {
	query := getproductstock.GetProductStockQuery{
		ProductId: entities.ProductID(req.Id),
	}

	product, err := s.getHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	return &productstockpb.GetProductStockResponse{
		Product: toProductStock(product),
	}, nil
}

func (s *ProductStockService) ReserveStock(
	ctx context.Context,
	req *productstockpb.ReserveStockRequest,
) (*productstockpb.ReserveStockResponse, error) {
	cmd := reserveproductstock.ReserveProductStockCommand{
		Products: toProductReserveStockCmd(req.Products),
	}

	if err := s.reserveHandler.Handle(ctx, cmd); err != nil {
		return nil, err
	}

	return &productstockpb.ReserveStockResponse{
		Status: true,
	}, nil
}

func (s *ProductStockService) Restock(
	ctx context.Context,
	req *productstockpb.RestockRequest,
) (*productstockpb.RestockResponse, error) {
	cmd := restockproductstock.RestockProductStockCommand{
		ProductID: entities.ProductID(req.Id),
		Quantity:  int(req.Quantity),
	}

	if err := s.restockHandler.Handle(ctx, cmd); err != nil {
		return &productstockpb.RestockResponse{
			Status: false,
		}, err
	}

	return &productstockpb.RestockResponse{
		Status: true,
	}, nil
}

func (s *ProductStockService) ReleaseStock(
	ctx context.Context,
	req *productstockpb.ReleaseStockRequest,
) (*emptypb.Empty, error) {
	cmd := releaseproductstock.ReleaseProductStockCommand{
		Products: toProductReleaseStockCmd(req.Products),
	}

	if err := s.releaseHandler.Handle(ctx, cmd); err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
