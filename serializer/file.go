package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToBinaryFIle writes protobuf msg to binary file
func WriteProtobufToBinaryFIle(msg proto.Message, filename string) error {
	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	if err = ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	return nil
}
