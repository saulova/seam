package interfaces

type TelemetryProviderInterface interface {
	Register(config interface{}) error
}
