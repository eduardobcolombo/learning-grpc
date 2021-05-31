module github.com/eduardobcolombo/learning-grpc/server

// local replace
// replace github.com/eduardobcolombo/learning-grpc/portpb => /Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/portpb
// replace github.com/eduardobcolombo/learning-grpc/server => /Users/eduardocolombo/go/src/github.com/eduardobcolombo/learning-grpc/server

go 1.16

require (
	github.com/eduardobcolombo/learning-grpc/portpb v0.0.0-00010101000000-000000000000
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/lib/pq v1.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/grpc v1.38.0
)
