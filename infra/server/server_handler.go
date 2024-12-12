package server

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

type ServerHandler struct {
	dependencyContainer *dependencies.DependencyContainer
	serverConfig        *ServerConfig
	app                 *fiber.App
	logger              interfaces.LoggerInterface
}

const ServerHandlerId = "infra.api.server.ServerHandler"
const HealthCheckLiveId = "health-check.live"
const HealthCheckReadyId = "health-check.ready"

func NewServerHandler() *ServerHandler {
	dependencyContainer := dependencies.GetDependencyContainer()

	serverConfig := dependencyContainer.GetDependency(ServerConfigId).(*ServerConfig)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	app := fiber.New(fiber.Config{
		ReadTimeout:                  serverConfig.ReadTimeout,
		AppName:                      serverConfig.AppName,
		Prefork:                      serverConfig.Prefork,
		ServerHeader:                 serverConfig.ServerHeader,
		StrictRouting:                serverConfig.StrictRouting,
		CaseSensitive:                serverConfig.CaseSensitive,
		UnescapePath:                 serverConfig.UnescapePath,
		ETag:                         serverConfig.ETag,
		BodyLimit:                    serverConfig.BodyLimit,
		Concurrency:                  serverConfig.Concurrency,
		WriteTimeout:                 serverConfig.WriteTimeout,
		IdleTimeout:                  serverConfig.IdleTimeout,
		ReadBufferSize:               serverConfig.ReadBufferSize,
		WriteBufferSize:              serverConfig.WriteBufferSize,
		CompressedFileSuffix:         serverConfig.CompressedFileSuffix,
		ProxyHeader:                  serverConfig.ProxyHeader,
		GETOnly:                      serverConfig.GETOnly,
		DisableKeepalive:             serverConfig.DisableKeepalive,
		DisableDefaultDate:           serverConfig.DisableDefaultDate,
		DisableDefaultContentType:    serverConfig.DisableDefaultContentType,
		DisableHeaderNormalizing:     serverConfig.DisableHeaderNormalizing,
		DisableStartupMessage:        serverConfig.DisableStartupMessage,
		StreamRequestBody:            serverConfig.StreamRequestBody,
		DisablePreParseMultipartForm: serverConfig.DisablePreParseMultipartForm,
		ReduceMemoryUsage:            serverConfig.ReduceMemoryUsage,
		Network:                      serverConfig.Network,
		EnableTrustedProxyCheck:      serverConfig.EnableTrustedProxyCheck,
		TrustedProxies:               serverConfig.TrustedProxies,
		EnableIPValidation:           serverConfig.EnableIPValidation,
		EnablePrintRoutes:            serverConfig.EnablePrintRoutes,
		EnableSplittingOnParsers:     serverConfig.EnableSplittingOnParsers,
	})

	instance := &ServerHandler{
		dependencyContainer: dependencyContainer,
		serverConfig:        serverConfig,
		app:                 app,
		logger:              logger,
	}

	dependencyContainer.AddDependency(HealthCheckLiveId, false)
	dependencyContainer.AddDependency(HealthCheckReadyId, false)
	dependencyContainer.AddDependency(ServerHandlerId, instance)

	return instance
}

func (s *ServerHandler) addHealthCheck() {
	if s.serverConfig.DisableHealthCheck {
		return
	}

	s.app.Use(
		healthcheck.New(healthcheck.Config{
			LivenessProbe: func(c *fiber.Ctx) bool {
				return s.dependencyContainer.GetDependency(HealthCheckLiveId).(bool)
			},
			LivenessEndpoint: s.serverConfig.HealthCheckLiveRoute,
			ReadinessProbe: func(c *fiber.Ctx) bool {
				return s.dependencyContainer.GetDependency(HealthCheckReadyId).(bool)
			},
			ReadinessEndpoint: s.serverConfig.HealthCheckReadyRoute,
		}),
	)
}

func (s *ServerHandler) GetApp() *fiber.App {
	return s.app
}

func (s *ServerHandler) StartServer() {
	s.addHealthCheck()

	var err error

	s.dependencyContainer.AddDependency(HealthCheckLiveId, true)

	if s.serverConfig.TLS {
		err = s.app.ListenTLS(s.serverConfig.Address, s.serverConfig.CertFile, s.serverConfig.KeyFile)
	} else {
		err = s.app.Listen(s.serverConfig.Address)
	}

	if err != nil {
		s.logger.Fatal("Server is not running!", err)

		return
	}
}

func (s *ServerHandler) SetHealthCheckReady(status bool) {
	s.dependencyContainer.AddDependency(HealthCheckReadyId, status)
}
