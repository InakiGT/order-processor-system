package main

import (
	"log"
	"net"

	"github.com/InakiGT/processor/inventory-service/src/api/grpc/gateways"
	"github.com/InakiGT/processor/inventory-service/src/internal/infra/persistence/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error while trying to load .env")
	}

	db := gorm.NewDBConnection()

	listener, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal("Error while trying to run the server: ", err.Error())
	}

	grpcServer := grpc.NewServer()

	gateways.RegisterProductStockService(grpcServer, db)

	log.Println("gRPC server listening on port :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
