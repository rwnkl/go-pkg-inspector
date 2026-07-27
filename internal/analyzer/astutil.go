package analyzer

import (
	"bytes"
	"go/format"
	"go/token"
	"go/ast"

	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

// ExprString returns formatted Go source for an expression.
func ExprString(fset *token.FileSet, expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	var buf bytes.Buffer
	if err := format.Node(&buf, fset, expr); err != nil {
		return "<invalid>"
	}
	return buf.String()
}

// Position converts a token.Pos into semantic.Position.
func Position(fset *token.FileSet, pos token.Pos) semantic.Position {
	p := fset.Position(pos)
	return semantic.Position{
		File:   p.Filename,
		Line:   p.Line,
		Column: p.Column,
	}
}