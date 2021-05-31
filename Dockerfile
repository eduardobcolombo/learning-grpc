FROM golang:1.16 as go-client

ARG LOCAL_PATH=/Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc
WORKDIR ${LOCAL_PATH}
COPY . .
WORKDIR ${LOCAL_PATH}/client
RUN go mod download
RUN go build -o client-app .

CMD ["/Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/client/client-app"]

FROM golang:1.16 as go-server

ARG LOCAL_PATH=/Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc
WORKDIR ${LOCAL_PATH}
COPY . .
WORKDIR ${LOCAL_PATH}/server
RUN go mod download
RUN go build -o server-app .

CMD ["/Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/server/server-app"]