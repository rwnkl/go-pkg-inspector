package analyzer

import (
	"github.com/rwnkl/go-package-inspector/internal/semantic"
	"github.com/rwnkl/go-package-inspector/internal/workspace"
)

type PackageContext struct {
	Workspace *workspace.Workspace

	WorkspacePackage *workspace.Package

	SemanticPackage *semantic.Package
}