proto:
	protoc --proto_path=./pkg/pb/proto ./pkg/pb/proto/*.proto  --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=. --go_out=.
server:
	go run cmd/main.go
