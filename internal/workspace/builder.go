package workspace

import (
	"go/ast"
	"path/filepath"
	// "go/token"
	// "go/types"

	"golang.org/x/tools/go/packages"
)

// Builder constructs an immutable Workspace.
type Builder struct {
	workspace *Workspace
}

// NewBuilder creates a new builder.
func NewBuilder() *Builder {
	return &Builder{
		workspace: &Workspace{},
	}
}

// SetModule stores module information.
func (b *Builder) SetModule(name, path string) {

	b.workspace.Module = &Module{
		Name: name,
		Path: path,
	}
}

// AddPackage converts a packages.Package into our model.
func (b *Builder) AddPackage(pkg *packages.Package) {

	p := &Package{
		Name: pkg.Name,
		Path: pkg.PkgPath,
		Pkg:  pkg,
	}

	if pkg.Types != nil {
		p.Types = pkg.Types
	}

	if pkg.Fset != nil {
		p.FileSet = pkg.Fset
	}

	if len(pkg.Syntax) > 0 {
		p.AST = make([]*ast.File, len(pkg.Syntax))
		copy(p.AST, pkg.Syntax)
	}

	for _, file := range pkg.GoFiles {

		p.Files = append(
			p.Files,
			&File{
				Name: filepath.Base(file),
				Path: file,
			},
		)
	}

	b.workspace.Packages =
		append(b.workspace.Packages, p)
}

// Workspace finalizes and returns the workspace.
func (b *Builder) Workspace() *Workspace {

	// Build indexes before returning.
	b.workspace.buildIndex()

	return b.workspace
}