package interfaces

type ActionsManagerInterface interface {
	LoadAction(actionId string, actionPluginDependencyId string, config interface{}) error
	GetAction(actionId string) interface{}
}
