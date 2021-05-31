package interfaces

import (
	"log"

	"github.com/academiadaweb/learning-grpc/server/application"
	"github.com/academiadaweb/learning-grpc/server/domain/entity"
)

type Port struct {
	portApp application.PortAppInterface
}

func NewPort(pApp application.PortAppInterface) *Port {
	return &Port{
		portApp: pApp,
	}
}

func (p *Port) SavePort() {
	// TODO, get the data from GRPC request
	var port = entity.Port{
		Name:        "asdf",
		City:        "dfgh",
		Country:     "adsf",
		Alias:       []interface{}{},
		Regions:     []interface{}{},
		Coordinates: []float64{},
		Province:    "asdf",
		Timezone:    "asd",
		Unlocs:      []string{},
		Code:        "aaaa",
	}

	_, saveErr := p.portApp.SavePort(&port)
	if saveErr != nil {
		log.Fatalf("Error %v", saveErr)
	}
}
