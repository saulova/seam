package filesystem

import (
	"fmt"

	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/routes"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type RouteSchema struct {
	GatewayPath string   `yaml:"gatewayPath"`
	Methods     []string `yaml:"methods"`
	Middlewares []string `yaml:"middlewares"`
	Action      string   `yaml:"action"`
}

type RoutesFileSchema struct {
	Routes map[string]RouteSchema `yaml:"routes"`
}

type RoutesRepository struct {
	configFileHandler *ConfigFileHandler
	cachedRoutes      map[string]*routes.RouteEntity
}

const RoutesRepositoryId = "infra.filesystem.RoutesRepository"
const RoutesConfigFileHandlerId = "infra.filesystem.RoutesConfigFileHandler"

func NewRoutesRepository() interfaces.RoutesProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(RoutesConfigFileHandlerId).(*ConfigFileHandler)

	instance := &RoutesRepository{
		configFileHandler: configFileHandler,
	}

	dependencyContainer.AddDependency(RoutesRepositoryId, instance)

	return instance
}

func (r *RoutesRepository) ListRoutes() (*dtos.ListRoutesOutput, error) {
	if len(r.cachedRoutes) > 0 {
		output := &dtos.ListRoutesOutput{
			Routes: r.cachedRoutes,
		}

		return output, nil
	}

	var routesFile RoutesFileSchema

	err := r.configFileHandler.Unmarshal(&routesFile)
	if err != nil {
		panic(err)
	}

	result := map[string]*routes.RouteEntity{}

	for key, routeConfig := range routesFile.Routes {
		route, err := routes.NewRouteEntity(&routes.RouteEntityInput{
			Id:             key,
			GatewayPath:    routeConfig.GatewayPath,
			Methods:        routeConfig.Methods,
			MiddlewaresIds: routeConfig.Middlewares,
			Action:         routeConfig.Action,
		})
		if err != nil {
			return nil, err
		}

		result[key] = route
	}

	output := &dtos.ListRoutesOutput{
		Routes: result,
	}

	r.cachedRoutes = result

	return output, nil
}

func (r *RoutesRepository) FindRouteById(id string) (*dtos.FindRouteByIdOutput, error) {
	listRoutesOutput, err := r.ListRoutes()
	if err != nil {
		return nil, err
	}

	route, ok := listRoutesOutput.Routes[id]
	if !ok {
		return nil, fmt.Errorf("route \"%s\" does not exist", id)
	}

	output := &dtos.FindRouteByIdOutput{
		Route: route,
	}

	return output, nil
}
