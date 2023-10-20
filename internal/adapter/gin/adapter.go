package gin

import (
	"github.com/fbriansyah/bank-ina-test/internal/application/port"
	"github.com/gin-gonic/gin"
)

type GinAdapter struct {
	service port.ServicePort
	router  *gin.Engine
}

type getUserByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
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
