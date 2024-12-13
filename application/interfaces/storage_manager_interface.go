package interfaces

type StoragesManagerInterface interface {
	HasStorageAdapter(storageId string) bool
	CreateStorageAdapter(storageId string, storagePluginDependencyId string, config interface{}) (interface{}, error)
	GetStorageAdapter(storageId string) interface{}
}
