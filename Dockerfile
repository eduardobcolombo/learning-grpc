FROM golang:1.16 as go-client

ARG LOCAL_PATH=/app
WORKDIR ${LOCAL_PATH}
COPY . .
WORKDIR ${LOCAL_PATH}/client
RUN go mod download
RUN go build -o client-app .

CMD ["/app/client/client-app"]

FROM golang:1.16 as go-server

ARG LOCAL_PATH=/app
WORKDIR ${LOCAL_PATH}
COPY . .
WORKDIR ${LOCAL_PATH}/server
RUN go mod download
RUN go build -o server-app .

CMD ["/app/server/server-app"]