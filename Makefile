##################
### Environment variables
##################
GRPC_HOST=localhost
GRPC_PORT=50053
API_PORT=8888
LOCAL_PATH=/app
DB_HOST=localhost
DB_PASSWORD=passwd
DB_USER=user
DB_NAME=db
DB_PORT=5432

## Docker image name
IMG := eduardobcolombo/grpc
VERSION := 1.0 # $(shell git rev-parse --short HEAD)
CLUSTER := grpc-cluster
NAMESPACE := grpc

##################
### Build section
##################
## Shortcut to build docker images.
build-img: build-server-img build-client-img
.PHONY: build-img
## Build a docker image for the server service.
build-server-img:
	@docker build \
	-f ./docker/Dockerfile.server \
	-t $(IMG)-server:$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.
.PHONY: build-server-img

## Build a docker image for the client service.
build-client-img:
	@docker build \
	-f ./docker/Dockerfile.client \
	-t $(IMG)-client:$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.
.PHONY: build-client-img

## Shortcut to build the binary files.
build-bin: build-server-bin build-client-bin
.PHONY: build-bin

## Build a binary file for the server service.
build-server-bin: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/server.bin ./cmd/server/main.go
.PHONY: build-server-bin

## Build a binary file for the client service.
build-client-bin: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:1.16 \
		go build -o bin/client.bin ./cmd/client/main.go
.PHONY: build-client-bin

##################
### KinD - Kubernetes in Docker section.
##################
## Shortcut to init KinD with the images.
kind-init: kind-create kind-config build-img kind-load
.PHONY: kind-init

## Creates a KinD cluster locally.
kind-create: 
	@kind create cluster --name $(CLUSTER)
.PHONY: kind-create

## Apply secrets based on the .env file
kind-apply-secrets:
	@kubectl create secret generic server-secrets \
		--from-env-file=./deploy/server.env \
		--namespace $(NAMESPACE) --cluster kind-$(CLUSTER) --dry-run=client -o yaml \
	> ./k8s/server-secrets.yaml
	@kubectl apply -f ./k8s/server-secrets.yaml --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
	@kubectl create secret generic client-secrets \
		--from-env-file=./deploy/client.env \
		--namespace $(NAMESPACE) --cluster kind-$(CLUSTER) --dry-run=client -o yaml \
	> ./k8s/client-secrets.yaml
	@kubectl apply -f ./k8s/client-secrets.yaml --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-apply-secrets

## Set the current cluster to the context and load env variables.
kind-config:
	@kubectl config set-context --current --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-config

## Load docker built images to the KinD cluster.
kind-load:
	@kind load docker-image $(IMG)-server:$(VERSION) --name $(CLUSTER)
	@kind load docker-image $(IMG)-client:$(VERSION) --name $(CLUSTER)
.PHONY: kind-load

## Apply the k8s folder with yaml files to the cluster.
kind-apply: 
	@kubectl apply -f ./k8s/namespace --cluster kind-$(CLUSTER)
	@kubectl apply -f ./k8s/db --cluster kind-$(CLUSTER)
	@kubectl apply -f ./k8s/deployment --cluster kind-$(CLUSTER)
.PHONY: kind-apply

# Delete all resources hostedn in the ./k8s
kind-delete:
	@kubectl delete -f ./k8s/namespace --cluster kind-$(CLUSTER)
.PHONY: kind-delete

# Delete the cluster.
kind-clean:
	@kind delete cluster --name $(CLUSTER)
.PHONY: kind-delete

# Monitoring with watch
kind-watch:
	@watch kubectl get all --namespace $(NAMESPACE)
.PHONY: kind-watch

db:
	kubectl exec -it pod/postgres-5d8fd768d6-qbbt4 --namespace $(NAMESPACE) --cluster kind-$(CLUSTER) -- \
	psql -h $(DB_HOST) -U $(DB_USER) --password -p $(DB_PORT) $(DB_NAME)
.PHONY: db



##################
### Docker Compose section
##################
## Spin up the docker compose file with the env variables set.
up: 
	@docker-compose \
	-f docker-compose.yml \
	--env-file ./deploy/docker-local.env \
	up --force-recreate --build -d
.PHONY: up

## Remove all the created containers based on the docker-compose.yaml.
clean: 
	@docker-compose \
	-f docker-compose.yml \
	--env-file ./deploy/docker-local.env \
	down --remove-orphans -v
.PHONY: clean

##################
### Golang tools section
##################
## Shortcut to get packages and put them in the vendor folder.
tidy:
	@go mod tidy
	@go mod vendor
.PHONY: tidy

## Run the server based on the environment variables.
run-server:
	@export GRPC_HOST=$(GRPC_HOST) \
		    GRPC_PORT=$(GRPC_PORT) \
		    DB_HOST=$(DB_HOST) \
		    DB_PASSWORD=$(DB_PASSWORD) \
		    DB_USER=$(DB_USER) \
		    DB_NAME=$(DB_NAME) \
		    DB_PORT=$(DB_PORT) \
			; \
	go run ./cmd/server/main.go
.PHONY: run-server

## Run the client based on the environment variables.
run-client:
	@export GRPC_HOST=$(GRPC_HOST) \
			GRPC_PORT=$(GRPC_PORT) \
			API_PORT=$(API_PORT) \
			; \
	go run ./cmd/client/main.go
.PHONY: run-client

## Run the lint to make sure all is good.
lint:
	@golangci-lint run
.PHONY: lint

## Generate the protobuf files.
generate:
	@go generate ./...
.PHONY: generate

## Run tests.
test:
	@echo "-> Testing"; \
	export HOST=$(GRPC_HOST) \
		   PORT=$(GRPC_PORT) \
		   API_PORT=$(API_PORT) \
		   DB_HOST=$(DB_HOST) \
		   DB_PASSWORD=$(DB_PASSWORD) \
		   DB_USER=$(DB_USER) \
		   DB_NAME=$(DB_NAME) \
		   DB_PORT=$(DB_PORT) \
		   ; \
	go test ./cmd/... -v -cover -count=1	
	go test ./internal/... -v -cover -count=1	
.PHONY: test