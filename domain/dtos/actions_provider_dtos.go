package dtos

import (
	"github.com/saulova/seam/domain/entities/actions"
)

type ListActionsOutput struct {
	Actions map[string]*actions.ActionEntity
}

type FindActionByIdOutput struct {
	Action *actions.ActionEntity
}
