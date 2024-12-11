package ws

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/wsfw/configs"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	gorilla "github.com/gorilla/websocket"
)

type WebSocketForward struct {
	logger interfaces.LoggerInterface
}

type WebSocketMessage struct {
	MessageType int
	Message     []byte
}

const WebSocketForwardId = "plugins.wsfw.ws.WebSocketForward"

func NewWebSocketForward() *WebSocketForward {
	dependencyContainer := dependencies.GetDependencyContainer()

	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	instance := &WebSocketForward{
		logger: logger,
	}

	dependencyContainer.AddDependency(WebSocketForwardId, instance)

	return instance
}

func (w *WebSocketForward) clientToUpstreamRoutine(clientConnection *websocket.Conn, clientToUpstreamChannel chan WebSocketMessage, inactivityHandler *InactivityHandler) {
	go func() {
		defer close(clientToUpstreamChannel)
		for {
			messageType, message, err := clientConnection.ReadMessage()
			if err != nil {
				w.logger.Error("client read error", err)

				return
			}

			inactivityHandler.ResetInactivityTimer()

			clientToUpstreamChannel <- struct {
				MessageType int
				Message     []byte
			}{MessageType: messageType, Message: message}
		}
	}()
}

func (w *WebSocketForward) upstreamToClientRoutine(upstreamConnection *gorilla.Conn, upstreamToClientChannel chan WebSocketMessage, inactivityHandler *InactivityHandler) {
	go func() {
		defer close(upstreamToClientChannel)
		for {
			messageType, message, err := upstreamConnection.ReadMessage()
			if err != nil {
				w.logger.Error("upstream read error", err)
				return
			}

			inactivityHandler.ResetInactivityTimer()

			upstreamToClientChannel <- struct {
				MessageType int
				Message     []byte
			}{MessageType: messageType, Message: message}
		}
	}()
}

func (w *WebSocketForward) proxyDataBetweenChannelsLoop(clientConnection *websocket.Conn, upstreamConnection *gorilla.Conn, clientToUpstreamChannel chan WebSocketMessage, upstreamToClientChannel chan WebSocketMessage) {
	for {
		select {
		case msgStruct, ok := <-clientToUpstreamChannel:
			if !ok {
				return
			}
			err := upstreamConnection.WriteMessage(msgStruct.MessageType, msgStruct.Message)
			if err != nil {
				w.logger.Error("error writing to upstream", err)
				return
			}

		case msgStruct, ok := <-upstreamToClientChannel:
			if !ok {
				return
			}
			err := clientConnection.WriteMessage(msgStruct.MessageType, msgStruct.Message)
			if err != nil {
				w.logger.Error("error writing to client", err)
				return
			}
		}
	}
}

func (w *WebSocketForward) GetForwardAction(config *configs.WebSocketActionConfig, url string) func(*fiber.Ctx) error {
	return websocket.New(func(clientConnection *websocket.Conn) {
		inactivityHandler := NewInactivityHandler(config)
		inactivityHandler.StartCheckInactivity()

		clientPingPongHandler := NewPingPongHandler(config, w.logger)
		clientPingPongHandler.SetPingPongHandlers(clientConnection)
		clientPingPongHandler.PingRoutine(clientConnection)

		upstreamConnection, _, err := gorilla.DefaultDialer.Dial(url, nil)
		if err != nil {
			w.logger.Error("error connecting to upstream", err)
			return
		}
		defer upstreamConnection.Close()

		clientToUpstreamChannel := make(chan WebSocketMessage)
		upstreamToClientChannel := make(chan WebSocketMessage)

		w.clientToUpstreamRoutine(clientConnection, clientToUpstreamChannel, inactivityHandler)
		w.upstreamToClientRoutine(upstreamConnection, upstreamToClientChannel, inactivityHandler)
		w.proxyDataBetweenChannelsLoop(clientConnection, upstreamConnection, clientToUpstreamChannel, upstreamToClientChannel)
	}, websocket.Config{
		HandshakeTimeout:  config.HandshakeTimeout,
		Subprotocols:      config.Subprotocols,
		Origins:           config.Origins,
		ReadBufferSize:    config.ReadBufferSize,
		WriteBufferSize:   config.WriteBufferSize,
		EnableCompression: config.EnableCompression,
	})
}
