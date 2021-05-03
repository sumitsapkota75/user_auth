package bootstrap

import (
	"context"
	"user/api/controllers"
	"user/api/repository"
	"user/api/routes"
	"user/api/services"
	"user/infrastructure"

	"go.uber.org/fx"
)

// Module exported for initializing application
var Module = fx.Options(
	controllers.Module,
	routes.Module,
	infrastructure.Module,
	services.Module,
	repository.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.RequestHandler,
	routes routes.Routes,
	env infrastructure.Env,
	logger infrastructure.Logger,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting Application")
			logger.Zap.Info("----------------------------")
			logger.Zap.Info("------- USER AUTH SYSTEM -------")
			logger.Zap.Info("----------------------------")
			go func() {
				routes.Setup()
				if env.ServerPort == "" {
					handler.Gin.Run()
				} else {
					handler.Gin.Run(env.ServerPort)
				}
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Zap.Info("Stopping Application")
			return nil
		},
	})
}
