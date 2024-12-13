package filesystem

import (
	"fmt"

	"github.com/saulova/seam/domain/dtos"
	"github.com/saulova/seam/domain/entities/storages"
	"github.com/saulova/seam/domain/interfaces"
	"github.com/saulova/seam/libs/dependencies"
)

type StorageSchema struct {
	Use    string      `yaml:"use"`
	Config interface{} `yaml:"config"`
}

type StoragesFileSchema struct {
	Storages map[string]StorageSchema `yaml:"storages"`
}

type StoragesRepository struct {
	configFileHandler *ConfigFileHandler
	cachedStorages    map[string]*storages.StorageEntity
}

const StoragesRepositoryId = "infra.filesystem.StoragesRepository"
const StoragesConfigFileHandlerId = "infra.filesystem.StoragesConfigFileHandler"

func NewStoragesRepository() interfaces.StoragesProviderInterface {
	dependencyContainer := dependencies.GetDependencyContainer()

	configFileHandler := dependencyContainer.GetDependency(StoragesConfigFileHandlerId).(*ConfigFileHandler)

	instance := &StoragesRepository{
		configFileHandler: configFileHandler,
		cachedStorages:    map[string]*storages.StorageEntity{},
	}

	dependencyContainer.AddDependency(StoragesRepositoryId, instance)

	return instance
}

func (m *StoragesRepository) ListStorages() (*dtos.ListStoragesOutput, error) {
	if len(m.cachedStorages) > 0 {
		output := &dtos.ListStoragesOutput{
			Storages: m.cachedStorages,
		}

		return output, nil
	}

	var storagesFile StoragesFileSchema

	err := m.configFileHandler.Unmarshal(&storagesFile)
	if err != nil {
		return nil, err
	}

	result := map[string]*storages.StorageEntity{}

	for key, storageConfig := range storagesFile.Storages {
		storage, err := storages.NewStorageEntity(&storages.StorageEntityInput{
			Id:     key,
			Use:    storageConfig.Use,
			Config: storageConfig.Config,
		})
		if err != nil {
			return nil, err
		}

		result[key] = storage
	}

	output := &dtos.ListStoragesOutput{
		Storages: result,
	}

	m.cachedStorages = result

	return output, nil
}

func (m *StoragesRepository) FindStorageById(id string) (*dtos.FindStorageByIdOutput, error) {
	listStoragesOutput, err := m.ListStorages()
	if err != nil {
		return nil, err
	}

	storage, ok := listStoragesOutput.Storages[id]
	if !ok {
		return nil, fmt.Errorf("storage \"%s\" does not exist", id)
	}

	output := &dtos.FindStorageByIdOutput{
		Storage: storage,
	}

	return output, nil
}
