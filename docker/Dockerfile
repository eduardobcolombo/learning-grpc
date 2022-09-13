FROM golang:1.16 as client-builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .
RUN go build -o /client-app /app/cmd/client/main.go

# CMD ["/client-app"]

FROM golang:1.16 as server-builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app
COPY . .
RUN go build -o /server-app /app/cmd/server/main.go

# CMD ["/server-app"]


FROM scratch as client

COPY --from=client-builder /client-app /svc

EXPOSE 8000
CMD ["/svc"]


FROM scratch as server

COPY --from=server-builder /server-app /svc

EXPOSE 8000
CMD ["/svc"]