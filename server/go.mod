module github.com/academiadaweb/learning-grpc/server

// local replace
replace github.com/academiadaweb/learning-grpc/portpb => /Users/eduardocolombo/go/src/github.com/academiadaweb/learning-grpc/portpb

replace github.com/academiadaweb/learning-grpc/server => /Users/eduardocolombo/go/src/github.com/academiadaweb/learning-grpc/server

go 1.16

require (
	github.com/academiadaweb/learning-grpc/portpb v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.38.0 // indirect
)
