package storages

import (
	"errors"
	"fmt"
	"strings"
)

type StorageEntity struct {
	Id     string
	Use    string
	Config interface{}
}

func validate(input *StorageEntityInput) error {
	if input.Id == "" {
		return errors.New("invalid storage id")
	}

	if input.Use == "" || (strings.HasPrefix("plugins.", input.Use) && strings.HasSuffix("Storage", input.Use)) {
		return fmt.Errorf("invalid use storage: '%s'", input.Use)
	}

	return nil
}

func NewStorageEntity(input *StorageEntityInput) (*StorageEntity, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	instance := &StorageEntity{
		Id:     input.Id,
		Use:    input.Use,
		Config: input.Config,
	}

	return instance, nil
}
