generate_proto:
	protoc --go_out=./ --go-grpc_out=./ --proto_path=./ proto/bookcabin.proto
	protoc --grpc-gateway_out=./ --proto_path=./ proto/bookcabin.proto

