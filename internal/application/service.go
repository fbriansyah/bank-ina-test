package application

import "github.com/fbriansyah/bank-ina-test/internal/application/port"

type Service struct {
	db         port.DatabasePort
	tokenMaker port.TokenMakerPort
}

func NewService(db port.DatabasePort, tokenMaker port.TokenMakerPort) *Service {
	return &Service{
		db:         db,
		tokenMaker: tokenMaker,
	}
}
