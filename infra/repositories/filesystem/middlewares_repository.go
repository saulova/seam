package filesystem

import (
	"fmt"

	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/middlewares"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type MiddlewareSchema struct {
	Use    string      `yaml:"use"`
	Config interface{} `yaml:"config"`
}

type MiddlewaresFileSchema struct {
	Middlewares       map[string]MiddlewareSchema `yaml:"middlewares"`
	GlobalMiddlewares []string                    `yaml:"globalMiddlewares"`
}

type MiddlewaresRepository struct {
	configFileHandler       *ConfigFileHandler
	cachedMiddlewares       map[string]*middlewares.MiddlewareEntity
	cachedGlobalMiddlewares []string
}

const MiddlewaresRepositoryId = "infra.filesystem.MiddlewaresRepository"
const MiddlewaresConfigFileHandlerId = "infra.filesystem.MiddlewaresConfigFileHandler"

func NewMiddlewaresRepository() interfaces.MiddlewaresProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(MiddlewaresConfigFileHandlerId).(*ConfigFileHandler)

	instance := &MiddlewaresRepository{
		configFileHandler: configFileHandler,
		cachedMiddlewares: map[string]*middlewares.MiddlewareEntity{},
	}

	dependencyContainer.AddDependency(MiddlewaresRepositoryId, instance)

	return instance
}

func (m *MiddlewaresRepository) ListMiddlewares() (*dtos.ListMiddlewaresOutput, error) {
	if len(m.cachedMiddlewares) > 0 {
		output := &dtos.ListMiddlewaresOutput{
			Middlewares: m.cachedMiddlewares,
		}

		return output, nil
	}

	var middlewaresFile MiddlewaresFileSchema

	err := m.configFileHandler.Unmarshal(&middlewaresFile)
	if err != nil {
		panic(err)
	}

	result := map[string]*middlewares.MiddlewareEntity{}

	for key, middlewareConfig := range middlewaresFile.Middlewares {
		middleware, err := middlewares.NewMiddlewareEntity(&middlewares.MiddlewareEntityInput{
			Id:     key,
			Use:    middlewareConfig.Use,
			Config: middlewareConfig.Config,
		})
		if err != nil {
			return nil, err
		}

		result[key] = middleware
	}

	output := &dtos.ListMiddlewaresOutput{
		Middlewares: result,
	}

	m.cachedMiddlewares = result

	return output, nil
}

func (m *MiddlewaresRepository) ListGlobalMiddlewares() (*dtos.ListGlobalMiddlewaresOutput, error) {
	if len(m.cachedGlobalMiddlewares) > 0 {
		output := &dtos.ListGlobalMiddlewaresOutput{
			GlobalMiddlewares: m.cachedGlobalMiddlewares,
		}

		return output, nil
	}

	var middlewaresFile MiddlewaresFileSchema

	err := m.configFileHandler.Unmarshal(&middlewaresFile)
	if err != nil {
		panic(err)
	}

	output := &dtos.ListGlobalMiddlewaresOutput{
		GlobalMiddlewares: middlewaresFile.GlobalMiddlewares,
	}

	m.cachedGlobalMiddlewares = middlewaresFile.GlobalMiddlewares

	return output, nil
}

func (m *MiddlewaresRepository) FindMiddlewareById(id string) (*dtos.FindMiddlewareByIdOutput, error) {
	listMiddlewaresOutput, err := m.ListMiddlewares()
	if err != nil {
		return nil, err
	}

	middleware, ok := listMiddlewaresOutput.Middlewares[id]
	if !ok {
		return nil, fmt.Errorf("middleware \"%s\" does not exist", id)
	}

	output := &dtos.FindMiddlewareByIdOutput{
		Middleware: middleware,
	}

	return output, nil
}
