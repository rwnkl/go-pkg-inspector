package analyzer

import (
	"go/ast"
)

type Walker struct {
	ctx *Context
}

func NewWalker(ctx *Context) *Walker {
	return &Walker{
		ctx: ctx,
	}
}

func (w *Walker) Files(fn func(*ast.File)) {

	for _, file := range w.ctx.Package.WorkspacePackage.AST {
		fn(file)
	}
}