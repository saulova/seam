package interfaces

type MiddlewaresManagerInterface interface {
	LoadMiddleware(middlewareId string, middlewarePluginDependencyId string, config interface{}) error
	RegisterGlobalMiddleware(middlewareId string) error
	GetMiddleware(middlewareId string) interface{}
}
