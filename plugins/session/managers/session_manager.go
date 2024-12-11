package managers

import (
	"encoding/gob"
	"errors"
	"time"

	"github.com/saulova/seam/libs/dependencies"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type SessionManager struct {
}

const SessionManagerId = "plugins.session.handlers.SessionManager"

const DateTimeFormat = "yyyy-MM-dd'T'HH:mm:ss.SSSZ"

func NewSessionManager() *SessionManager {
	dependencyContainer := dependencies.GetDependencyContainer()

	gob.Register(map[string]interface{}{})

	instance := &SessionManager{}

	dependencyContainer.AddDependency(SessionManagerId, instance)

	return instance
}

func (s *SessionManager) getSessionFromLocals(ctx *fiber.Ctx) (map[string]interface{}, error) {
	sessionDataMap := ctx.Locals("session")
	if sessionDataMap == nil {
		return nil, errors.New("missing session")
	}

	return sessionDataMap.(map[string]interface{}), nil
}

func (s *SessionManager) destroySession(sessionStore *session.Store, ctx *fiber.Ctx) error {
	sessionInstance, err := sessionStore.Get(ctx)
	if err != nil {
		return err
	}

	sessionInstance.Destroy()

	return nil
}

func (s *SessionManager) renewSession(sessionInstance *session.Session, sessionDataMap map[string]interface{}, autoRenewAfter time.Duration) (bool, error) {
	isAfter := true

	autoRenewAfterDateTime, ok := sessionDataMap["autoRenewAfterDateTime"]
	if !ok {
		return false, nil
	}

	renewAfter, err := time.Parse(time.RFC3339, autoRenewAfterDateTime.(string))
	if err != nil {
		return false, err
	}

	isAfter = time.Now().After(renewAfter)

	if !isAfter {
		return false, nil
	}

	err = sessionInstance.Regenerate()
	if err != nil {
		return false, err
	}

	sessionDataMap["autoRenewAfterDateTime"] = time.Now().Add(autoRenewAfter).UTC().Format(time.RFC3339)

	return true, nil
}

func (s *SessionManager) saveSession(sessionStore *session.Store, ctx *fiber.Ctx) error {
	sessionDataMap, err := s.getSessionFromLocals(ctx)
	if err != nil {
		return err
	}

	sessionInstance, err := sessionStore.Get(ctx)
	if err != nil {
		return err
	}

	for _, key := range []string{"destroy", "save"} {
		delete(sessionDataMap, key)
	}

	sessionInstance.Set("session", sessionDataMap)

	err = sessionInstance.Save()
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionManager) LoadSession(sessionStore *session.Store, autoRenewAfter time.Duration, ctx *fiber.Ctx) error {
	sessionInstance, err := sessionStore.Get(ctx)
	if err != nil {
		return err
	}

	sessionValue := sessionInstance.Get("session")

	sessionDataMap := make(map[string]interface{}, 0)

	if sessionValue != nil {
		sessionDataMap = sessionValue.(map[string]interface{})
	}

	isNewSession, err := s.renewSession(sessionInstance, sessionDataMap, autoRenewAfter)
	if err != nil {
		return err
	}

	if isNewSession {
		s.saveSession(sessionStore, ctx)
	}

	sessionDataMap["destroy"] = func() {
		s.destroySession(sessionStore, ctx)
	}

	sessionDataMap["save"] = func() {
		s.saveSession(sessionStore, ctx)
	}

	ctx.Locals("session", sessionDataMap)

	return nil
}
