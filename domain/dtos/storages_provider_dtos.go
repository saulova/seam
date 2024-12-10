package dtos

import (
	"github.com/saulova/seam/domain/entities/storages"
)

type ListStoragesOutput struct {
	Storages map[string]*storages.StorageEntity
}

type FindStorageByIdOutput struct {
	Storage *storages.StorageEntity
}
