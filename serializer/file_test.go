package serializer_test

import (
	"testing"

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
}
