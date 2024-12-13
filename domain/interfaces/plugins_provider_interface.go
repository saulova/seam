package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type PluginsProviderInterface interface {
	ListPlugins() (*dtos.ListPluginsOutput, error)
}
