//go:generate 	protoc --go_out=./ --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative	./ports.proto
//go:generate   mockgen --destination=./mock_port/mock_ports_grpc.pb.go --source=./ports_grpc.pb.go

package portpb
