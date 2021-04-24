package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/petrostrak/gRPC-with-Go/pb"
	"github.com/petrostrak/gRPC-with-Go/sample"
	"github.com/petrostrak/gRPC-with-Go/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFIle := "../tmp/laptop.bin"
	laptop1 := sample.NewLaptop()
	if err := serializer.WriteProtobufToBinaryFIle(laptop1, binaryFIle); err != nil {
		require.NoError(t, err)
	}

	laptop2 := &pb.Laptop{}
	if err := serializer.ReadProtobufFromBinaryFile(binaryFIle, laptop2); err != nil {
		require.NoError(t, err)
		require.True(t, proto.Equal(laptop1, laptop2))
	}
}
