package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/csrf/configs"

	"github.com/gofiber/fiber/v2"
	csrfMw "github.com/gofiber/fiber/v2/middleware/csrf"
)

type CSRFMiddleware struct {
	storagesMediator interfaces.StoragesMediatorInterface
}

const CSRFMiddlewareId = "plugins.csrf.middlewares.CSRFMiddleware"

func NewCSRFMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	storagesMediator := dependencyContainer.GetDependency(interfaces.StoragesMediatorInterfaceId).(interfaces.StoragesMediatorInterface)

	instance := &CSRFMiddleware{
		storagesMediator: storagesMediator,
	}

	dependencyContainer.AddDependency(CSRFMiddlewareId, instance)

	return instance
}

func (c *CSRFMiddleware) createCSRFConfig(config *configs.CSRFMiddlewareConfig) (csrfMw.Config, error) {
	var storage fiber.Storage = nil

	if config.Storage != "" {
		storageAdapter, err := c.storagesMediator.GetStorage(config.Storage)
		if err != nil {
			panic(err)
		}

		storage = storageAdapter.(fiber.Storage)
	}

	CSRFConfig := csrfMw.Config{
		KeyLookup:         config.KeyLookup,
		CookieName:        config.CookieName,
		CookieDomain:      config.CookieDomain,
		CookiePath:        config.CookiePath,
		CookieSecure:      config.CookieSecure,
		CookieHTTPOnly:    config.CookieHTTPOnly,
		CookieSameSite:    config.CookieSameSite,
		CookieSessionOnly: config.CookieSessionOnly,
		Expiration:        config.CookieExpiration,
		SingleUseToken:    config.SingleUseToken,
		Storage:           storage,
	}

	return CSRFConfig, nil
}

func (c *CSRFMiddleware) Build(config interface{}) (interface{}, error) {
	csrfMiddlewareConfig, err := configs.NewCSRFMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	csrfConfig, err := c.createCSRFConfig(csrfMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := csrfMw.New(csrfConfig)

	return middlewareFunc, nil
}
