package commands

import (
	applicationInterfaces "github.com/saulova/seam/application/interfaces"
	domainInterfaces "github.com/saulova/seam/domain/interfaces"
	libsInterfaces "github.com/saulova/seam/libs/interfaces"
)

type GetStorageAdapterCommand struct {
	storagesProvider domainInterfaces.StoragesProviderInterface
	storagesManager  applicationInterfaces.StoragesManagerInterface
	logger           libsInterfaces.LoggerInterface
}

func NewGetStorageAdapterCommand(storagesProvider domainInterfaces.StoragesProviderInterface, storagesManager applicationInterfaces.StoragesManagerInterface, logger libsInterfaces.LoggerInterface) *GetStorageAdapterCommand {
	return &GetStorageAdapterCommand{
		storagesProvider: storagesProvider,
		storagesManager:  storagesManager,
		logger:           logger,
	}
}

func (g *GetStorageAdapterCommand) Execute(storageId string) (interface{}, error) {
	g.logger.Debug("getting storage", "storageId", storageId)

	if !g.storagesManager.HasStorageAdapter(storageId) {
		g.logger.Debug("creating storage", "storageId", storageId)

		findStorageByIdOutput, err := g.storagesProvider.FindStorageById(storageId)
		if err != nil {
			g.logger.Error("find storage error", err)
			return nil, err
		}

		g.logger.Debug("storage found by id", "output", findStorageByIdOutput)

		storage, err := g.storagesManager.CreateStorageAdapter(findStorageByIdOutput.Storage.Id, findStorageByIdOutput.Storage.Use, findStorageByIdOutput.Storage.Config)
		if err != nil {
			g.logger.Error("create storage error", err)
			return nil, err
		}

		g.logger.Debug("storage created")

		return storage, nil
	}

	storage := g.storagesManager.GetStorageAdapter(storageId)

	g.logger.Info("storage found", "storageId", storageId)

	return storage, nil
}
