package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/petrostrak/gRPC-with-Go/pb/pb"
	"github.com/petrostrak/gRPC-with-Go/sample"
	"github.com/petrostrak/gRPC-with-Go/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()
	if err := serializer.WriteProtobufToBinaryFIle(laptop1, binaryFile); err != nil {
		require.NoError(t, err)
	}

	laptop2 := &pb.Laptop{}
	if err := serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2); err != nil {
		require.NoError(t, err)
		require.True(t, proto.Equal(laptop1, laptop2))
	}

	if err := serializer.WriteProtobufToJSON(laptop1, jsonFile); err != nil {
		require.NoError(t, err)
	}
}
