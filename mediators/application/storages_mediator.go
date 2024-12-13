package application

import (
	"github.com/saulova/seam/infra/managers"
	"github.com/saulova/seam/infra/repositories/filesystem"

	"github.com/saulova/seam/application/commands"
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type StoragesMediator struct {
	getStorageCommand *commands.GetStorageAdapterCommand
	logger            interfaces.LoggerInterface
}

const StoragesMediatorId = interfaces.StoragesMediatorInterfaceId

func NewStoragesMediator() interfaces.StoragesMediatorInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	storagesRepository := dependencyContainer.GetDependency(filesystem.StoragesRepositoryId).(*filesystem.StoragesRepository)
	storagesManager := dependencyContainer.GetDependency(managers.StoragesManagerId).(*managers.StoragesManager)
	logger := dependencyContainer.GetDependency(interfaces.LoggerInterfaceId).(interfaces.LoggerInterface)

	getStorageAdapterCommand := commands.NewGetStorageAdapterCommand(storagesRepository, storagesManager, logger)

	instance := &StoragesMediator{
		getStorageCommand: getStorageAdapterCommand,
		logger:            logger,
	}

	dependencyContainer.AddDependency(StoragesMediatorId, instance)

	return instance
}

func (s *StoragesMediator) GetStorage(storageId string) (interface{}, error) {
	s.logger.Debug("executing storages mediator get storage")

	storage, err := s.getStorageCommand.Execute(storageId)
	if err != nil {
		return nil, err
	}

	s.logger.Debug("storages mediator get storage executed")

	return storage, nil
}
