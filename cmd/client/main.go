package main

import (
	"fmt"
	"os"

	"github.com/eduardobcolombo/learning-grpc/cmd/client/api"
	"github.com/eduardobcolombo/learning-grpc/internal/pkg/foundation"
)

func main() {

	// Construct the application logger.
	log, err := foundation.NewLogger("CLIENT")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	code := api.Initialize(log)
	os.Exit(code)
}
