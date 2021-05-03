package routes

import (
	"user/infrastructure"

	"github.com/gin-gonic/gin"
)

//UserRoutes struct
type UserRoutes struct {
	logger  infrastructure.Logger
	handler infrastructure.RequestHandler
}

// Setup user Routes
func (s UserRoutes) Setup(c *gin.Context) {
	s.logger.Zap.Info("Setting up user routes...")
	user := s.handler.Gin.Group("/users")
	{
		user.GET("", gin.H{"data": "data"})
	}
}

func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.RequestHandler,
) UserRoutes {
	return UserRoutes{
		handler: handler,
		logger:  logger,
	}
}
