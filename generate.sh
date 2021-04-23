!#bin/bash
protoc --proto_path=proto proto/*.proto --go-grpc_out=. --go_out=. proto/*.proto