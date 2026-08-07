package main

import (
	"log"
	"net"
	"os"

	"github.com/InakiGT/processor/order-service/src/api/grpc/gateways"
	persistence "github.com/InakiGT/processor/order-service/src/internal/infra/persistence/gorm"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error while trying to load .env file. Using env variables instead. ", err)
	}

	port := os.Getenv("PORT")

	db := persistence.NewDBConnection()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error while trying to start the server: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	grpcClient, err := grpc.NewClient(
		os.Getenv("ORDER_SERVICE"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("Error while trying to create a new client: ", err.Error())
	}

	gateways.RegisterOrderService(grpcServer, db, grpcClient)

	log.Println("Server running on port ", port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Error while trying to attach gRPC service: ", err.Error())
	}
}
