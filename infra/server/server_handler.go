package server

import (
	"github.com/saulova/seam/libs/dependencies"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
)

type ServerHandler struct {
	dependencyContainer *dependencies.DependencyContainer
	serverConfig        *ServerConfig
	app                 *fiber.App
}

const ServerHandlerId = "infra.api.server.ServerHandler"

func NewServerHandler() *ServerHandler {
	dependencyContainer := dependencies.GetDependencyContainer()

	serverConfig := dependencyContainer.GetDependency(ServerConfigId).(*ServerConfig)

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
	}

	dependencyContainer.AddDependency("health-check.live", false)
	dependencyContainer.AddDependency("health-check.ready", false)
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
				return s.dependencyContainer.GetDependency("health-check.live").(bool)
			},
			LivenessEndpoint: s.serverConfig.HealthCheckLiveRoute,
			ReadinessProbe: func(c *fiber.Ctx) bool {
				return s.dependencyContainer.GetDependency("health-check.ready").(bool)
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

	s.dependencyContainer.AddDependency("health-check.live", true)

	if s.serverConfig.TLS {
		err = s.app.ListenTLS(s.serverConfig.Address, s.serverConfig.CertFile, s.serverConfig.KeyFile)
	} else {
		err = s.app.Listen(s.serverConfig.Address)
	}

	if err != nil {
		log.Fatal("Oops... Server is not running! Reason: %v", err)

		return
	}
}

func (s *ServerHandler) SetHealthCheckReady(status bool) {
	s.dependencyContainer.AddDependency("health-check.ready", status)
}
