package managers

import (
	"fmt"

	"github.com/saulova/seam/application/interfaces"
	"github.com/saulova/seam/infra/server"
	"github.com/saulova/seam/libs/dependencies"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"

	"github.com/gofiber/fiber/v2"
)

type MiddlewaresManager struct {
	dependencyContainer *dependencies.DependencyContainer
}

const MiddlewaresManagerId = "infra.middlewares.MiddlewaresManager"

func NewMiddlewaresManager() interfaces.MiddlewaresManagerInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &MiddlewaresManager{
		dependencyContainer: dependencyContainer,
	}

	dependencyContainer.AddDependency(MiddlewaresManagerId, instance)

	return instance
}

func (m *MiddlewaresManager) normalizeMiddlewareFuncDependencyId(middlewareId string) string {
	return fmt.Sprintf("middlewares.%s", middlewareId)
}

func (m *MiddlewaresManager) LoadMiddleware(middlewareId string, middlewarePluginDependencyId string, config interface{}) error {
	middlewareFuncDependencyId := m.normalizeMiddlewareFuncDependencyId(middlewareId)

	middlewarePlugin := m.dependencyContainer.GetDependency(middlewarePluginDependencyId).(libsInterfaces.BuilderInterface)

	middlewareFunc, err := middlewarePlugin.Build(config)
	if err != nil {
		return err
	}

	m.dependencyContainer.AddDependency(middlewareFuncDependencyId, middlewareFunc.(func(ctx *fiber.Ctx) error))

	return nil
}

func (m *MiddlewaresManager) RegisterGlobalMiddleware(middlewareId string) error {
	globalMiddlewareRegister := server.NewGlobalMiddlewareRegister()
	middleware := m.GetMiddleware(middlewareId).(func(ctx *fiber.Ctx) error)

	globalMiddlewareRegister.UseMiddleware(middleware)
	globalMiddlewareRegister.Register()

	return nil
}

func (m *MiddlewaresManager) GetMiddleware(middlewareId string) interface{} {
	middlewareFuncDependencyId := m.normalizeMiddlewareFuncDependencyId(middlewareId)

	middlewareAdapter := m.dependencyContainer.GetDependency(middlewareFuncDependencyId)

	return middlewareAdapter
}
