package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *GinAdapter) setupRouter() {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "OPTION", "POST", "PUT", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Authorization"}

	router.Use(cors.New(corsConfig))

	router.POST("/login", s.Login)

	router.POST("/users", s.Register)
	router.GET("/users", s.ListUsers)
	router.GET("/users/:id", s.GetUserByID)
	router.PUT("/users/:id", s.UpdateUser)

	s.router = router
}
