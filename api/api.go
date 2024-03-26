package api

import (
	"github.com/gin-gonic/gin"
	"management-personal/config"
)

type Service struct {
	*gin.Engine
}

func NewService() *Service {
	return &Service{
		gin.Default(),
	}
}

func (s *Service) Start() {
	apiPort := config.Env.ApiPort

	s.GetRoutes()
	s.Engine.Run(apiPort)
}
