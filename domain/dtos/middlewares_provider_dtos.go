package dtos

import (
	"github.com/saulova/seam/domain/entities/middlewares"
)

type ListMiddlewaresOutput struct {
	Middlewares map[string]*middlewares.MiddlewareEntity
}

type ListGlobalMiddlewaresOutput struct {
	GlobalMiddlewares []string
}

type FindMiddlewareByIdOutput struct {
	Middleware *middlewares.MiddlewareEntity
}
