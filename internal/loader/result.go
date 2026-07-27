package loader

import (
	"github.com/rwnkl/go-package-inspector/internal/workspace"
)

type Result struct {

	Workspace *workspace.Workspace

	Diagnostics []Diagnostic
}