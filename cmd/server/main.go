package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/petrostrak/gRPC-with-Go/pb"
	"github.com/petrostrak/gRPC-with-Go/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")

	laptopServer := service.NewLaptopServer(laptopStore, imageStore)
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tpc", address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
