GRPC_HOST=localhost
GRPC_PORT=50053
API_PORT=8888
LOCAL_PATH=/app
DB_HOST=localhost
DB_PASSWORD=passwd
DB_USER=user
DB_NAME=db
DB_PORT=5432

lint:
	@golangci-lint run
.PHONY: lint

generate:
	@protoc --go_out=./internal/pkg/portpb/ --go_opt=paths=source_relative --go-grpc_out=./internal/pkg/portpb/ --go-grpc_opt=paths=source_relative ./internal/pkg/portpb/ports.proto

build:
	@docker-compose -f docker-compose.yml --env-file ./deploy/app.env build --no-cache --force-rm

up: 
	@docker-compose -f docker-compose.yml --env-file ./deploy/app.env up --force-recreate -d

clean: 
	@docker-compose -f docker-compose.yml --env-file ./deploy/app.env down --remove-orphans -v

tidy:
	@go mod tidy
	@go mod vendor

run-server:
	export GRPC_HOST=$(GRPC_HOST); \
	export GRPC_PORT=$(GRPC_PORT); \
	export DB_HOST=$(DB_HOST); \
	export DB_PASSWORD=$(DB_PASSWORD); \
	export DB_USER=$(DB_USER); \
	export DB_NAME=$(DB_NAME); \
	export DB_PORT=$(DB_PORT); \
	go run ./cmd/server/main.go

run-client:
	export GRPC_HOST=$(GRPC_HOST); \
	export GRPC_PORT=$(GRPC_PORT); \
	export API_PORT=$(API_PORT); \
	go run ./cmd/client/main.go

build-server: 
# docker-compose -f docker-compose.yml --env-file ./deploy/api.env build --no-cache --force-rm --pull
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/server.bin ./cmd/server/main.go

build-client: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/client.bin ./cmd/client/main.go		

.PHONY: test
test:
	@echo "-> Testing"
	export HOST=$(GRPC_HOST); \
	export PORT=$(GRPC_PORT); \
	export API_PORT=$(API_PORT); \
	export DB_HOST=$(DB_HOST); \
	export DB_PASSWORD=$(DB_PASSWORD); \
	export DB_USER=$(DB_USER); \
	export DB_NAME=$(DB_NAME); \
	export DB_PORT=$(DB_PORT); \
	go test ./cmd/... -v -cover -count=1	
	go test ./internal/... -v -cover -count=1	

