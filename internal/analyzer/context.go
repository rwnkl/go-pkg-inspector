package analyzer

import (
	"github.com/rwnkl/go-package-inspector/internal/semantic"
	"github.com/rwnkl/go-package-inspector/internal/workspace"
)

type Context struct {
	Workspace *workspace.Workspace

	Model *semantic.Model

	Package *PackageContext

	Table *semantic.SymbolTable
}

func (c *Context) Register(s semantic.Symbol) {
	c.Table.Add(s)
}