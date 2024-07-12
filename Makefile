gen :
	protoc --go_out=. --go-grpc_out=. proto/process_manager.proto proto/memory_message.proto proto/keyboard_message.proto proto/laptop_message.proto proto/screen_message.proto proto/storage_message.proto proto/laptop_service.proto proto/filter_message.proto

clean :
	rm -rf psm
 
prisma-generate:
	go run github.com/steebchen/prisma-client-go generate dev

migrate :
	go run github.com/steebchen/prisma-client-go migrate dev
# run :
# 	go build
# 	./grpc

run :
	go run main.go

server : 
	go run cmd/server/main.go -port 8080

client :
	go run cmd/client/main.go -address 0.0.0.0:8080
test:
	rm -rf temp/*
	go test -v ./...