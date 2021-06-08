GRPC_HOST = localhost
GRPC_PORT = 50053
API_PORT = 8888
LOCAL_PATH = /app
DB_HOST=localhost
DB_PASSWORD=passwd
DB_USER=user
DB_NAME=db
DB_PORT=5433

generate:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portpb/ports.proto

up: 
	@docker compose up -d

down: 
	@docker compose down

run_s:
	@cd server; \
	export PORT=$(GRPC_PORT); \
	export DB_HOST=$(DB_HOST); \
	export DB_PASSWORD=$(DB_PASSWORD); \
	export DB_USER=$(DB_USER); \
	export DB_NAME=$(DB_NAME); \
	export DB_PORT=$(DB_PORT); \
	go run main.go

run_c:
	@echo "-> Running Client"
	@cd client; \
	export HOST=$(GRPC_HOST); \
	export PORT=$(GRPC_PORT); \
	export API_PORT=$(API_PORT); \
	go run main.go

build_s: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH)/server -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/server.bin main.go

build_c: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH)/client -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/client.bin main.go		

test_c:
	@echo "-> Testing Client"
	@cd client; \
	export HOST=$(GRPC_HOST); \
	export PORT=$(GRPC_PORT); \
	export API_PORT=$(API_PORT); \
	go test -v -cover ./...

test_s:
	@echo "-> Testing Server"
	@cd server; \
	export PORT=$(GRPC_PORT); \
	export DB_HOST=$(DB_HOST); \
	export DB_PASSWORD=$(DB_PASSWORD); \
	export DB_USER=$(DB_USER); \
	export DB_NAME=$(DB_NAME); \
	export DB_PORT=$(DB_PORT); \
	go test -v -cover ./...
