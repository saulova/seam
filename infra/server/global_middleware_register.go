package server

import (
	"github.com/saulova/seam/libs/dependencies"

	"github.com/gofiber/fiber/v2"
)

type GlobalMiddlewareRegister struct {
	app        *fiber.App
	middleware func(*fiber.Ctx) error
}

func NewGlobalMiddlewareRegister() *GlobalMiddlewareRegister {
	dependencyContainer := dependencies.GetDependencyContainer()

	serverHandler := dependencyContainer.GetDependency(ServerHandlerId).(*ServerHandler)

	instance := &GlobalMiddlewareRegister{
		app:        serverHandler.GetApp(),
		middleware: nil,
	}

	return instance
}

func (a *GlobalMiddlewareRegister) UseMiddleware(middleware func(*fiber.Ctx) error) {
	a.middleware = middleware
}

func (a *GlobalMiddlewareRegister) Register() {
	a.app.Use(a.middleware)
}
