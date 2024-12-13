package managers

import (
	"fmt"

	"github.com/saulova/seam/application/interfaces"
	"github.com/saulova/seam/libs/dependencies"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type StoragesManager struct {
	dependencyContainer *dependencies.DependencyContainer
}

const StoragesManagerId = "infra.storages.StoragesManager"

func NewStoragesManager() interfaces.StoragesManagerInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &StoragesManager{
		dependencyContainer: dependencyContainer,
	}

	dependencyContainer.AddDependency(StoragesManagerId, instance)

	return instance
}

func (s *StoragesManager) normalizeStorageAdapterDependencyId(storageId string) string {
	return fmt.Sprintf("storages.%s", storageId)
}

func (s *StoragesManager) HasStorageAdapter(storageId string) bool {
	storageAdapterDependencyId := s.normalizeStorageAdapterDependencyId(storageId)

	return s.dependencyContainer.HasDependency(storageAdapterDependencyId)
}

func (s *StoragesManager) CreateStorageAdapter(storageId string, storagePluginDependencyId string, config interface{}) (interface{}, error) {
	storageAdapterDependencyId := s.normalizeStorageAdapterDependencyId(storageId)

	storagePlugin := s.dependencyContainer.GetDependency(storagePluginDependencyId).(libsInterfaces.BuilderInterface)

	storageAdapter, err := storagePlugin.Build(config)
	if err != nil {
		return nil, err
	}

	s.dependencyContainer.AddDependency(storageAdapterDependencyId, storageAdapter)

	return storageAdapter, nil
}

func (s *StoragesManager) GetStorageAdapter(storageId string) interface{} {
	storageAdapterDependencyId := s.normalizeStorageAdapterDependencyId(storageId)

	storageAdapter := s.dependencyContainer.GetDependency(storageAdapterDependencyId)

	return storageAdapter
}
