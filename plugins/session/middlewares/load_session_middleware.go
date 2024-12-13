package middlewares

import (
	"strings"

	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/session/configs"
	"github.com/saulova/seam/plugins/session/managers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type LoadSessionMiddleware struct {
	sessionManager   *managers.SessionManager
	storagesMediator interfaces.StoragesMediatorInterface
	logger           interfaces.LoggerInterface
}

const LoadSessionMiddlewareId = "plugins.session.middlewares.LoadSessionMiddleware"

func NewLoadSessionMiddleware() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	sessionManager := dependencyContainer.GetDependency(managers.SessionManagerId).(*managers.SessionManager)
	storagesMediator := dependencyContainer.GetDependency(interfaces.StoragesMediatorInterfaceId).(interfaces.StoragesMediatorInterface)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	instance := &LoadSessionMiddleware{
		sessionManager:   sessionManager,
		storagesMediator: storagesMediator,
		logger:           logger,
	}

	dependencyContainer.AddDependency(LoadSessionMiddlewareId, instance)

	return instance
}

func (l *LoadSessionMiddleware) getStorage(storageId string) fiber.Storage {
	var storage fiber.Storage = nil

	if storageId != "" {
		storageAdapter, err := l.storagesMediator.GetStorage(storageId)
		if err != nil {
			panic(err)
		}

		storage = storageAdapter.(fiber.Storage)
	}

	return storage
}

func (l *LoadSessionMiddleware) Build(config interface{}) (interface{}, error) {
	loadSessionMiddlewareConfig, err := configs.NewLoadSessionMiddlewareConfig(config)
	if err != nil {
		return nil, err
	}

	sessionStore := session.New(session.Config{
		Storage:           l.getStorage(loadSessionMiddlewareConfig.Storage),
		KeyLookup:         loadSessionMiddlewareConfig.KeyLookup,
		CookieHTTPOnly:    loadSessionMiddlewareConfig.CookieHTTPOnly,
		CookieSecure:      loadSessionMiddlewareConfig.CookieSecure,
		CookieSameSite:    loadSessionMiddlewareConfig.CookieSameSite,
		CookieSessionOnly: loadSessionMiddlewareConfig.CookieSessionOnly,
		Expiration:        loadSessionMiddlewareConfig.CookieExpiration,
	})

	middlewareFunc := func(ctx *fiber.Ctx) error {
		err := l.sessionManager.LoadSession(sessionStore, loadSessionMiddlewareConfig.AutoRenewAfter, ctx)
		if err != nil {
			l.logger.Error("load session error", err)

			return ctx.Next()
		}

		if loadSessionMiddlewareConfig.DisableSessionForward {
			ctx.Request().Header.DelCookie(strings.Split(loadSessionMiddlewareConfig.KeyLookup, ":")[1])
		}

		return ctx.Next()
	}

	return middlewareFunc, nil
}
