package middlewares

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/limiter/configs"

	"github.com/gofiber/fiber/v2"
	limiterMw "github.com/gofiber/fiber/v2/middleware/limiter"
)

type LimiterMiddleware struct {
	storagesMediator interfaces.StoragesMediatorInterface
}

const LimiterMiddlewareId = "plugins.limiter.middlewares.LimiterMiddleware"

func NewLimiterMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	storagesMediator := dependencyContainer.GetDependency(interfaces.StoragesMediatorInterfaceId).(interfaces.StoragesMediatorInterface)

	instance := &LimiterMiddleware{
		storagesMediator: storagesMediator,
	}

	dependencyContainer.AddDependency(LimiterMiddlewareId, instance)

	return instance
}

func (l *LimiterMiddleware) createLimiterConfig(config *configs.LimiterMiddlewareConfig) (limiterMw.Config, error) {
	var storage fiber.Storage = nil

	if config.Storage != "" {
		storageAdapter, err := l.storagesMediator.GetStorage(config.Storage)
		if err != nil {
			panic(err)
		}

		storage = storageAdapter.(fiber.Storage)
	}

	var limiterAlgorithm limiterMw.LimiterHandler = nil

	if config.SlidingWindow {
		limiterAlgorithm = limiterMw.SlidingWindow{}
	}

	limiterConfig := limiterMw.Config{
		Max:                    config.MaxConnections,
		Expiration:             config.Expiration,
		SkipFailedRequests:     config.SkipFailedRequests,
		SkipSuccessfulRequests: config.SkipSuccessfulRequests,
		Storage:                storage,
		LimiterMiddleware:      limiterAlgorithm,
	}

	return limiterConfig, nil
}

func (l *LimiterMiddleware) Build(config interface{}) (interface{}, error) {
	limiterMiddlewareConfig, err := configs.NewLimiterMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	limiterConfig, err := l.createLimiterConfig(limiterMiddlewareConfig)
	if err != nil {
		return nil, err
	}

	middlewareFunc := limiterMw.New(limiterConfig)

	return middlewareFunc, nil
}
