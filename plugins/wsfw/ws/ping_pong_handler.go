package ws

import (
	"net"
	"time"

	"github.com/saulova/seam/libs/interfaces"
	"github.com/saulova/seam/plugins/wsfw/configs"

	"github.com/gofiber/contrib/websocket"
)

type PingPongHandler struct {
	pongTimeout  time.Duration
	pingInterval time.Duration
	logger       interfaces.LoggerInterface
}

func NewPingPongHandler(config *configs.WebSocketActionConfig, logger interfaces.LoggerInterface) *PingPongHandler {
	return &PingPongHandler{
		pongTimeout:  config.PongTimeout,
		pingInterval: config.PingInterval,
		logger:       logger,
	}
}

func (p *PingPongHandler) SetPingPongHandlers(conn *websocket.Conn) {
	conn.SetPingHandler(func(appData string) error {
		p.logger.Debug("ping", "data", appData)

		err := conn.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(p.pongTimeout))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}

		return err
	})

	conn.SetPongHandler(func(appData string) error {
		p.logger.Debug("pong", "data", appData)

		return nil
	})
}

func (p *PingPongHandler) PingRoutine(conn *websocket.Conn) {
	go func() {
		ticker := time.NewTicker(p.pingInterval)
		defer ticker.Stop()

		for range ticker.C {
			if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				p.logger.Error("error while send ping to (", conn.RemoteAddr().String(), ")", err)

				return
			}
		}
	}()
}
