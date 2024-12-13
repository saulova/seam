package interfaces

import (
	"github.com/saulova/seam/libs/interfaces"
)

type PluginLoaderInterface interface {
	LoadPlugin(pluginPath string) (interfaces.PluginInterface, error)
}
