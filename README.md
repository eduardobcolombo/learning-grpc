# GRPC Learning

    This is a grpc test learning built in golang 1.16 using protobuf.


## Protobuf
To run the protobuf you can use the below command:
```protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative portpb/ports.proto ```

## Docker 
In order to make things easy for devs, I created a make file to run docker commands, so fell free to check the flavours:

`$ make generate` will generate the grpc files based on the protobuf messages

`$ make up` will start the docker-composer with all instances (Db, Client and Server).
You can test if it is working firing some curls like below:

`curl http://localhost:8888/v1/ports -X POST` It will load all ports to the DB
`curl http://localhost:8888/v1/ports -X GET` It will get all ports from DB

`$ make down` will put the docker-composer down

`$ make run-server` will run the local grpc server

`$ make run-client` will run the local grpc client to send message to the server

`$ make build-server` will build the server go package and export it to server/bin folder

`$ make build-client` will build the client go package and export it to client/bin folder

`$ make test` will run the test suite

## Next steps:

The ports.json file does not have enough data to have an unique identifier like ID or UUID 
or something like this.

Figure out a way to identify ports and allow it to be updated instead of just inserting.

Add more tests

Add more grpc status messages through the client server communication using codes and status from the packages
`"google.golang.org/grpc/codes"`
`"google.golang.org/grpc/status"`
