package storages

import (
	"github.com/saulova/seam/libs/dependencies"
	"github.com/saulova/seam/libs/interfaces"
)

type ValkeyStorage struct{}

const ValkeyStorageId = "plugins.valkey.storages.ValkeyStorage"

func NewValkeyStorage() interfaces.BuilderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	instance := &ValkeyStorage{}

	dependencyContainer.AddDependency(ValkeyStorageId, instance)

	return instance
}

func (v *ValkeyStorage) Build(config interface{}) (interface{}, error) {
	storage, err := NewValkeyStorageAdapter(config)
	if err != nil {
		return nil, err
	}

	return storage, nil
}
