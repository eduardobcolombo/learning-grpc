##################
### Environment variables
##################
include ./deploy/make.env

VERSION=1.0 # $(shell git rev-parse --short HEAD)

##################
### Build section
##################
## Shortcut to build the binary files.
build-bin: build-server-bin build-client-bin
.PHONY: build-bin

## Build a binary file for the server service.
build-server-bin: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:$(GO_VERSION) \
		go build -o bin/server.bin ./cmd/server/main.go
.PHONY: build-server-bin

## Build a binary file for the client service.
build-client-bin: 
	@docker run --rm \
		-v ${PWD}:$(LOCAL_PATH) \
		-w $(LOCAL_PATH) -e GOOS=linux -e GOARCH=amd64 \
		golang:$(GO_VERSION) \
		go build -o bin/client.bin ./cmd/client/main.go
.PHONY: build-client-bin

## Shortcut to build docker images.
build-img: build-server-img build-client-img
.PHONY: build-img
## Build a docker image for the server service.
build-server-img:
	docker build \
	-f ./deploy/docker/Dockerfile.server \
	-t $(IMG)server:$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.
.PHONY: build-server-img

## Build a docker image for the client service.
build-client-img:
	docker build \
	-f ./deploy/docker/Dockerfile.client \
	-t $(IMG)client:$(VERSION) \
	--build-arg BUILD_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.
.PHONY: build-client-img

##################
### Docker Compose section
##################
## Spin up the docker compose file with the env variables set.
up: 
	@docker compose \
	-f docker-compose.yml  \
	up --force-recreate --build -d
.PHONY: up

## Remove all the created containers based on the docker compose.yaml.
clean: 
	@docker compose \
	-f docker-compose.yml  \
	down --remove-orphans -v
.PHONY: clean


##################
### KinD - Kubernetes in Docker section.
##################
## Shortcut to init KinD with the images.
kind-init: kind-create kind-config build-img kind-load kind-apply kind-apply-secrets
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
	> ./deploy/k8s/server-secrets.yaml
	@kubectl apply -f ./deploy/k8s/server-secrets.yaml --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
	@kubectl create secret generic client-secrets \
		--from-env-file=./deploy/client.env \
		--namespace $(NAMESPACE) --cluster kind-$(CLUSTER) --dry-run=client -o yaml \
	> ./deploy/k8s/client-secrets.yaml
	@kubectl apply -f ./deploy/k8s/client-secrets.yaml --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-apply-secrets

## Set the current cluster to the context and load env variables.
kind-config:
	@kubectl config set-context --current --namespace $(NAMESPACE) --cluster kind-$(CLUSTER)
.PHONY: kind-config

## Load docker built images to the KinD cluster.
kind-load:
	@kind load docker-image $(IMG)server:$(VERSION) --name $(CLUSTER)
	@kind load docker-image $(IMG)client:$(VERSION) --name $(CLUSTER)
.PHONY: kind-load

## Apply the k8s folder with yaml files to the cluster.
kind-apply: 
	@kubectl apply -f ./deploy/k8s/namespace --cluster kind-$(CLUSTER)
	@kubectl apply -f ./deploy/k8s/db --cluster kind-$(CLUSTER)
	@kubectl apply -f ./deploy/k8s/deployment --cluster kind-$(CLUSTER)
	@kubectl apply -f ./deploy/k8s/traefik --cluster kind-$(CLUSTER)
.PHONY: kind-apply

# Delete all resources hostedn in the ./deploy/k8s
kind-delete:
	@kubectl delete -f ./deploy/k8s/namespace --cluster kind-$(CLUSTER)
.PHONY: kind-delete

# Delete the cluster.
kind-clean:
	@kind delete cluster --name $(CLUSTER)
.PHONY: kind-delete

# Monitoring with watch
kind-watch:
	@watch kubectl get all --namespace $(NAMESPACE)
.PHONY: kind-watch

expose:
	@kubectl port-forward service/client 8888
.PHONY: expose

# Get a DB session
db:
	kubectl exec -it deploy/postgres --namespace $(NAMESPACE) --cluster kind-$(CLUSTER) -- \
	psql -h $(DB_HOST) -U $(DB_USER) --password -p $(DB_PORT) $(DB_NAME)
.PHONY: db

##################
### Golang tools section
##################
## Shortcut to get packages and put them in the vendor folder.
tidy:
	@go mod tidy
	@go mod vendor
.PHONY: tidy

## Run the lint to make sure all is good.
lint:
	docker run --rm -t \
		-v "$(PWD):/app" \
		--workdir /app \
		golangci/golangci-lint:v1.62.2 golangci-lint run
.PHONY: lint

## Generate the protobuf files.
generate:
	@go generate ./...
.PHONY: generate

## Run tests.
test: 
	go test -v -mod=vendor -coverprofile=coverage.out -covermode=set -cover  -count=1 -timeout 60s \
	$$(go list ./... | grep -v '/testhelpers' | grep -v '/portpb')
	@coverage=$$(go tool cover -func coverage.out | tail -n 1 | awk '{print $$3;}'); \
		sed -i.bak -e "s/.*test coverage.*/$${coverage} test coverage./" README.md && rm README.md.bak
.PHONY: test

# Send the ports.json file to the client.
send: 
	@curl -v -F file=@data/ports.json http://localhost:8888/v1/ports
.PHONY: send

get: 
	@curl http://localhost:8888/v1/ports -X GET
.PHONY: get

healthcheck: liveness readiness

liveness: 
	@curl http://localhost:8888/liveness -X GET
.PHONY: get

readiness: 
	@curl http://localhost:8888/readiness -X GET
.PHONY: get