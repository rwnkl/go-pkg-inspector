package workspace

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/packages"
)

// Package contains everything we know about one Go package.
type Package struct {

	// Name of package.
	Name string

	// Import path.
	Path string

	// Loaded package.
	Pkg *packages.Package

	// Type information.
	Types *types.Package

	// FileSet used while parsing.
	FileSet *token.FileSet

	// Parsed AST.
	AST []*ast.File

	// Source files.
	Files []*File
}