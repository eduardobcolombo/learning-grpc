generate:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portpb/ports.proto

run_s:
	@cd server; \
	export PORT=:50051; \
	go run main.go

run_c:
	@echo "-> Running Client"
	@cd client; \
	export HOST=localhost; \
	export PORT=:50051; \
	go run main.go
