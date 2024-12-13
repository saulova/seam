package interfaces

const LoggerInterfaceId = "libs.interfaces.LoggerInterface"

type LoggerInterface interface {
	SetLevel(logLevel string)
	Debug(message string, keysAndValues ...interface{})
	Info(message string, keysAndValues ...interface{})
	Warn(message string, keysAndValues ...interface{})
	Error(message string, keysAndValues ...interface{})
	Fatal(message string, keysAndValues ...interface{})
}
