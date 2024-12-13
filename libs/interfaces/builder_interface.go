package interfaces

type BuilderInterface interface {
	Build(config interface{}) (interface{}, error)
}
