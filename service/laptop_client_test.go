package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/petrostrak/gRPC-with-Go/pb"
	"github.com/petrostrak/gRPC-with-Go/sample"
	"github.com/petrostrak/gRPC-with-Go/serializer"
	"github.com/petrostrak/gRPC-with-Go/service"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopStore := service.NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	laptop := sample.NewLaptop()
	expectedId := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedId, res.Id)

	// check that the laptop is saved to the store
	other, err := laptopStore.Find(res.Id)
	if err != nil {
		require.NoError(t, err)
		require.NotNil(t, other)
	}

	// check that the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)
}

func requireSameLaptop(t *testing.T, laptop1, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	require.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)
}

func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)

	return pb.NewLaptopServiceClient(conn)
}

func startTestLaptopServer(t *testing.T, laptopStore service.LaptopStore) string {
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)

	grpcServer.Serve(listener)

	return listener.Addr().String()
}
