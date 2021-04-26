package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/petrostrak/gRPC-with-Go/pb"
	"github.com/petrostrak/gRPC-with-Go/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()

	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	// add gRPC interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	// call evans with:
	// evans -r repl -p 8080
	reflection.Register(grpcServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tpc", address)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
