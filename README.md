# GRPC Learning

    This is a grpc test learning built in golang 1.16 using protobuf.


## Protobuf

To run the protobuf you can use the below command:
`$ make generate`

## Docker 

In order to make things easy for devs, I created a make file to run docker commands, so fell free to check the flavours:

`$ make up` will start the docker-composer with all instances (Db, Client and Server).
You can test if it is working firing some curls like below:

`$ make send` It will load all ports to the DB
`$ make get` It will get all ports from DB

`$ make down` will put the docker-composer down

## Kubernetes

We are running it using KinD to provide a local cluster.

`$ make kind-init` Shortcut to init KinD with the images.

`$ make kind-watch` Monitoring pods with watch.

`$ make db` Get a DB session.




`$ make kind-create` Creates a KinD cluster locally.

`$ make kind-apply-secrets` Apply secrets based on the .env file

`$ make kind-config` Set the current cluster to the context and load env variables.

`$ make kind-load` Load docker built images to the KinD cluster.

`$ make kind-apply` Apply the k8s folder with yaml files to the cluster.

`$ make kind-delete` Delete all resources hostedn in the ./k8s

`$ make kind-clean` Delete the cluster.

## Running locally

`$ make run-server` will run the local grpc server

`$ make run-client` will run the local grpc client to send message to the server

## Builing locally binary

`$ make build-server-bin` will build the server go package and export it to the bin folder

`$ make build-client-bin` will build the client go package and export it to the bin folder

## Running tests

`$ make test` will run the test suite

## Running lint

`$ make lint` will run the linter

## Useful tools

`$ make expose` will run expose the api to access locally.

## Next steps:

The ports.json file does not have enough data to have an unique identifier like ID or UUID 
or something like this.

Figure out a way to identify ports and allow it to be updated instead of just inserting.

Add more tests

Add more grpc status messages through the client server communication using codes and status from the packages
`"google.golang.org/grpc/codes"`
`"google.golang.org/grpc/status"`
