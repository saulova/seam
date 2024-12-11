package ws

import (
	"time"

	"github.com/saulova/seam/plugins/wsfw/configs"
)

type InactivityHandler struct {
	maxInactivity   time.Duration
	inactivityTimer *time.Timer
}

func NewInactivityHandler(config *configs.WebSocketActionConfig) *InactivityHandler {
	return &InactivityHandler{
		maxInactivity:   config.MaxInactivity,
		inactivityTimer: time.NewTimer(config.MaxInactivity),
	}
}

func (i *InactivityHandler) ResetInactivityTimer() {
	if !i.inactivityTimer.Stop() {
		<-i.inactivityTimer.C
	}

	i.inactivityTimer.Reset(i.maxInactivity)
}

func (i *InactivityHandler) StartCheckInactivity(closeFunctions ...func() error) {
	go func() {
		for range i.inactivityTimer.C {
			println("Conexão encerrada por inatividade")
			for _, closeFunction := range closeFunctions {
				closeFunction()
			}
		}
	}()
}
