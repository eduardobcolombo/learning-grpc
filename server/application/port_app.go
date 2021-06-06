package application

import (
	"github.com/eduardobcolombo/learning-grpc/server/domain/entity"
	"github.com/eduardobcolombo/learning-grpc/server/domain/repository"
)

type portApp struct {
	pr repository.PortRepository
}

var _ PortAppInterface = &portApp{}

type PortAppInterface interface {
	SavePort(*entity.Port) (*entity.Port, error)
	RetrievePorts() ([]*entity.Port, error)
}

func (p *portApp) SavePort(port *entity.Port) (*entity.Port, error) {
	return p.pr.SavePort(port)
}

func (p *portApp) RetrievePorts() ([]*entity.Port, error) {
	return p.pr.RetrievePorts()
}
