package dtos

import (
	"github.com/saulova/seam/domain/entities/plugins"
)

type ListPluginsOutput struct {
	Plugins []*plugins.PluginEntity
}
