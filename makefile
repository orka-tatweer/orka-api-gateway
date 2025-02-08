generate:
	protoc --go_out=. --go-grpc_out=. --twirp_out=. pkg/pb/service.proto

