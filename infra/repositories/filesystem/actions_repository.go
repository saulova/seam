package filesystem

import (
	"fmt"

	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/actions"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type ActionSchema struct {
	Use    string      `yaml:"use"`
	Config interface{} `yaml:"config"`
}

type ActionsFileSchema struct {
	Actions map[string]ActionSchema `yaml:"actions"`
}

type ActionsRepository struct {
	configFileHandler *ConfigFileHandler
	cachedActions     map[string]*actions.ActionEntity
}

const ActionsRepositoryId = "infra.filesystem.ActionsRepository"
const ActionsConfigFileHandlerId = "infra.filesystem.ActionsConfigFileHandler"

func NewActionsRepository() interfaces.ActionsProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(ActionsConfigFileHandlerId).(*ConfigFileHandler)

	instance := &ActionsRepository{
		configFileHandler: configFileHandler,
		cachedActions:     map[string]*actions.ActionEntity{},
	}

	dependencyContainer.AddDependency(ActionsRepositoryId, instance)

	return instance
}

func (m *ActionsRepository) ListActions() (*dtos.ListActionsOutput, error) {
	if len(m.cachedActions) > 0 {
		output := &dtos.ListActionsOutput{
			Actions: m.cachedActions,
		}

		return output, nil
	}

	var actionsFile ActionsFileSchema

	err := m.configFileHandler.Unmarshal(&actionsFile)
	if err != nil {
		panic(err)
	}

	result := map[string]*actions.ActionEntity{}

	for key, actionConfig := range actionsFile.Actions {
		action, err := actions.NewActionEntity(&actions.ActionEntityInput{
			Id:     key,
			Use:    actionConfig.Use,
			Config: actionConfig.Config,
		})
		if err != nil {
			return nil, err
		}

		result[key] = action
	}

	output := &dtos.ListActionsOutput{
		Actions: result,
	}

	m.cachedActions = result

	return output, nil
}

func (m *ActionsRepository) FindActionById(id string) (*dtos.FindActionByIdOutput, error) {
	listActionsOutput, err := m.ListActions()
	if err != nil {
		return nil, err
	}

	action, ok := listActionsOutput.Actions[id]
	if !ok {
		return nil, fmt.Errorf("action \"%s\" does not exist", id)
	}

	output := &dtos.FindActionByIdOutput{
		Action: action,
	}

	return output, nil
}
