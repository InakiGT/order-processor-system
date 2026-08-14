package main

import (
	"log"
	"os"

	"github.com/InakiGT/processor/api-gateway/src/api"
	"github.com/InakiGT/processor/api-gateway/src/api/http/inventory"
	"github.com/InakiGT/processor/api-gateway/src/api/http/order"
	orderpb "github.com/InakiGT/processor/api-gateway/src/api/pb/order/v1"
	productstockpb "github.com/InakiGT/processor/api-gateway/src/api/pb/product_stock/v1"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error while trying to load .env file. Using env variables instead")
	}

	orderConn, err := grpc.NewClient(
		os.Getenv("ORDER_SERVICE"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("An error ocurred while trying to create a grpc client:", err)
	}

	inventoryConn, err := grpc.NewClient(
		os.Getenv("INVENTORY_SERVICE"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	orderHandler := order.NewOrderHandler(
		orderpb.NewOrderServiceClient(orderConn),
	)

	inventoryHandler := inventory.NewInventoryHandler(
		productstockpb.NewProductStockServiceClient(inventoryConn),
	)

	router := api.NewRouter(
		orderHandler,
		inventoryHandler,
	)

	router.Run()
}
