gen:
	protoc --proto_path=proto proto/*.proto --go-grpc_out=. --go_out=. proto/*.proto

clean:
	rm proto/*pb.go

run:
	go run main.go