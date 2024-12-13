package server

import (
	"fmt"
	"strings"

	"github.com/saulova/seam/libs/dependencies"

	"github.com/gofiber/fiber/v2"
)

type RouteRegister struct {
	app             *fiber.App
	path            string
	methodFunctions []func(path string, handlers ...func(*fiber.Ctx) error) fiber.Router
	middlewares     []func(ctx *fiber.Ctx) error
	action          func(ctx *fiber.Ctx) error
}

func NewRouteRegister() *RouteRegister {
	dependencyContainer := dependencies.GetDependencyContainer()

	serverHandler := dependencyContainer.GetDependency(ServerHandlerId).(*ServerHandler)

	instance := &RouteRegister{
		app:             serverHandler.GetApp(),
		methodFunctions: []func(path string, handlers ...func(*fiber.Ctx) error) fiber.Router{},
		middlewares:     []func(ctx *fiber.Ctx) error{},
		action:          nil,
	}

	return instance
}

func (r *RouteRegister) UsePath(path string) {
	r.path = path
}

func (r *RouteRegister) UseMethods(methods []string) error {
	for _, method := range methods {
		switch strings.ToLower(method) {
		case "get":
			r.methodFunctions = append(r.methodFunctions, r.app.Get)
		case "post":
			r.methodFunctions = append(r.methodFunctions, r.app.Post)
		case "put":
			r.methodFunctions = append(r.methodFunctions, r.app.Put)
		case "delete":
			r.methodFunctions = append(r.methodFunctions, r.app.Delete)
		case "head":
			r.methodFunctions = append(r.methodFunctions, r.app.Head)
		case "options":
			r.methodFunctions = append(r.methodFunctions, r.app.Options)
		case "patch":
			r.methodFunctions = append(r.methodFunctions, r.app.Patch)
		case "all":
			r.methodFunctions = append(r.methodFunctions, r.app.All)
		default:
			return fmt.Errorf("method '%s' not exist", method)
		}
	}

	return nil
}

func (r *RouteRegister) UseMiddlewares(middlewares []func(*fiber.Ctx) error) {
	r.middlewares = middlewares
}

func (r *RouteRegister) UseAction(action func(*fiber.Ctx) error) {
	r.action = action
}

func (r *RouteRegister) Register() {
	for _, methodFunction := range r.methodFunctions {
		handlers := r.middlewares
		handlers = append(handlers, r.action)

		methodFunction(r.path, handlers...)
	}
}
