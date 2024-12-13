package filesystem

import (
	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/services"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type ServiceSchema struct {
	GatewayBasePath string   `yaml:"gatewayBasePath"`
	MiddlewaresIds  []string `yaml:"middlewares"`
	RoutesIds       []string `yaml:"routes"`
}

type ServicesFileSchema struct {
	Services map[string]ServiceSchema `yaml:"services"`
}

type ServicesRepository struct {
	configFileHandler *ConfigFileHandler
	cachedServices    map[string]*services.ServiceEntity
}

const ServicesRepositoryId = "infra.filesystem.ServicesRepository"
const ServicesConfigFileHandlerId = "infra.filesystem.ServicesConfigFileHandler"

func NewServicesRepository() interfaces.ServicesProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(ServicesConfigFileHandlerId).(*ConfigFileHandler)

	instance := &ServicesRepository{
		configFileHandler: configFileHandler,
	}

	dependencyContainer.AddDependency(ServicesRepositoryId, instance)

	return instance
}

func (s *ServicesRepository) ListServices() (*dtos.ListServicesOutput, error) {
	if len(s.cachedServices) > 0 {
		output := &dtos.ListServicesOutput{
			Services: s.cachedServices,
		}

		return output, nil
	}

	var servicesFile ServicesFileSchema

	err := s.configFileHandler.Unmarshal(&servicesFile)
	if err != nil {
		panic(err)
	}

	result := map[string]*services.ServiceEntity{}

	for key, serviceConfig := range servicesFile.Services {
		service, err := services.NewServiceEntity(&services.ServiceEntityInput{
			Id:              key,
			GatewayBasePath: serviceConfig.GatewayBasePath,
			MiddlewaresIds:  serviceConfig.MiddlewaresIds,
			RoutesIds:       serviceConfig.RoutesIds,
		})
		if err != nil {
			return nil, err
		}

		result[key] = service
	}

	output := &dtos.ListServicesOutput{
		Services: result,
	}

	s.cachedServices = result

	return output, nil
}
