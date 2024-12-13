package interfaces

import (
	"github.com/saulova/seam/domain/dtos"
)

type ActionsProviderInterface interface {
	ListActions() (*dtos.ListActionsOutput, error)
	FindActionById(id string) (*dtos.FindActionByIdOutput, error)
}
