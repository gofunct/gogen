package generator

import (
	"github.com/gofunct/common/ui"
)

type status int

const (
	statusCreate status = iota
	statusDelete
	statusExist
	statusIdentical
	statusConflicted
	statusForce
	statusSkipped
)

var (
	creatableStatusSet = map[status]struct{}{
		statusCreate: {},
		statusForce:  {},
	}
)

func (s status) Fprint(ui ui.UI, msg string) {
	switch s {
	case statusCreate, statusForce:
		ui.Success(msg)
	case statusDelete:
		ui.Warn(msg)
	case statusConflicted:
		ui.Error(msg)
	default:
		ui.Info(msg)
	}
}

func (s status) ShouldCreate() bool {
	_, ok := creatableStatusSet[s]
	return ok
}
