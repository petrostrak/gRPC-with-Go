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

// ReadProtobufFromBinaryFile reads protobuf msg from binary file
func ReadProtobufFromBinaryFile(filename string, msg proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	if err = proto.Unmarshal(data, msg); err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	return nil
}
