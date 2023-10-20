package gin

import (
	"github.com/fbriansyah/bank-ina-test/internal/application/port"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	service port.ServicePort
	router  *gin.Engine
}

func NewAdapter(service port.ServicePort) *GinAdapter {
	server := &GinAdapter{
		service: service,
	}

	server.setupRouter()

	return server
}

func (s *GinAdapter) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
