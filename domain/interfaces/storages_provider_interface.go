package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type StoragesProviderInterface interface {
	ListStorages() (*dtos.ListStoragesOutput, error)
	FindStorageById(id string) (*dtos.FindStorageByIdOutput, error)
}
