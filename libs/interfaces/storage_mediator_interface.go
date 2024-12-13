package interfaces

const StoragesMediatorInterfaceId = "libs.interfaces.StoragesMediatorInterface"

type StoragesMediatorInterface interface {
	GetStorage(storageId string) (interface{}, error)
}
