package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/helmet/configs"

	helmetMw "github.com/gofiber/fiber/v2/middleware/helmet"
)

type HelmetMiddleware struct{}

const HelmetMiddlewareId = "plugins.helmet.middlewares.HelmetMiddleware"

func NewHelmetMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &HelmetMiddleware{}

	dependencyContainer.AddDependency(HelmetMiddlewareId, instance)

	return instance
}

func (c *HelmetMiddleware) createHelmetConfig(config *configs.HelmetMiddlewareConfig) (helmetMw.Config, error) {
	helmetConfig := helmetMw.Config{
		XSSProtection:             config.XSSProtection,
		ContentTypeNosniff:        config.ContentTypeNosniff,
		XFrameOptions:             config.XFrameOptions,
		HSTSMaxAge:                config.HSTSMaxAge,
		HSTSExcludeSubdomains:     config.HSTSExcludeSubdomains,
		ContentSecurityPolicy:     config.ContentSecurityPolicy,
		CSPReportOnly:             config.CSPReportOnly,
		HSTSPreloadEnabled:        config.HSTSPreloadEnabled,
		ReferrerPolicy:            config.ReferrerPolicy,
		PermissionPolicy:          config.PermissionPolicy,
		CrossOriginEmbedderPolicy: config.CrossOriginEmbedderPolicy,
		CrossOriginOpenerPolicy:   config.CrossOriginOpenerPolicy,
		CrossOriginResourcePolicy: config.CrossOriginResourcePolicy,
		OriginAgentCluster:        config.OriginAgentCluster,
		XDNSPrefetchControl:       config.XDNSPrefetchControl,
		XDownloadOptions:          config.XDownloadOptions,
		XPermittedCrossDomain:     config.XPermittedCrossDomain,
	}

	return helmetConfig, nil
}

func (c *HelmetMiddleware) Build(config interface{}) (interface{}, error) {
	helmetMiddlewareConfig, err := configs.NewHelmetMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	helmetConfig, err := c.createHelmetConfig(helmetMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := helmetMw.New(helmetConfig)

	return middlewareFunc, nil
}
