package managers

import (
	"fmt"

	"github.com/saulova/seam/application/interfaces"
	"github.com/saulova/seam/libs/dependencies"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"

	"github.com/gofiber/fiber/v2"
)

type ActionsManager struct {
	dependencyContainer *dependencies.DependencyContainer
}

const ActionsManagerId = "infra.actions.ActionsManager"

func NewActionsManager() interfaces.ActionsManagerInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &ActionsManager{
		dependencyContainer: dependencyContainer,
	}

	dependencyContainer.AddDependency(ActionsManagerId, instance)

	return instance
}

func (m *ActionsManager) normalizeActionFuncDependencyId(actionId string) string {
	return fmt.Sprintf("actions.%s", actionId)
}

func (m *ActionsManager) LoadAction(actionId string, actionPluginDependencyId string, config interface{}) error {
	actionFuncDependencyId := m.normalizeActionFuncDependencyId(actionId)

	actionPlugin := m.dependencyContainer.GetDependency(actionPluginDependencyId).(libsInterfaces.BuilderInterface)

	actionFunc, err := actionPlugin.Build(config)
	if err != nil {
		return err
	}

	m.dependencyContainer.AddDependency(actionFuncDependencyId, actionFunc.(func(ctx *fiber.Ctx) error))

	return nil
}

func (m *ActionsManager) GetAction(actionId string) interface{} {
	actionFuncDependencyId := m.normalizeActionFuncDependencyId(actionId)

	actionAdapter := m.dependencyContainer.GetDependency(actionFuncDependencyId)

	return actionAdapter
}
