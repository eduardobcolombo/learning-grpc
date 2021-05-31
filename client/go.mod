module github.com/eduardobcolombo/learning-grpc/client

// local replace
// replace github.com/eduardobcolombo/learning-grpc/portpb => /Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/portpb
// replace github.com/eduardobcolombo/learning-grpc/client => /Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/client

go 1.16

require (
	github.com/eduardobcolombo/learning-grpc/portpb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.38.0
)
