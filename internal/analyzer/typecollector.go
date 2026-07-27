package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

type TypeCollector struct{}

func (*TypeCollector) Collect(
	ctx *Context,
) error {

	for _, pkg := range ctx.Workspace.Packages {

		sp := &semantic.Package{
			Name: pkg.Name,
			Path: pkg.Path,
		}

		ctx.Model.Packages =
			append(ctx.Model.Packages, sp)

		walker := NewWalker(ctx)
		
		walker.TypeSpecs(func(ts *ast.TypeSpec) {})

		for _, file := range pkg.AST {

			ast.Inspect(file, func(n ast.Node) bool {

				gen, ok := n.(*ast.GenDecl)

				if !ok || gen.Tok != token.TYPE {
					return true
				}

				for _, spec := range gen.Specs {

					ts := spec.(*ast.TypeSpec)

					t := &semantic.Type{
						// Name: ts.Name.Name,
					}

					switch ts.Type.(type) {

					case *ast.StructType:
						t.Kind = semantic.TypeStruct

					case *ast.InterfaceType:
						t.Kind = semantic.TypeInterface

					default:

						if ts.Assign.IsValid() {
							// t.Kind = semantic.TypeAlias
						} else {
							// t.Kind = semantic.TypeDefined
						}
					}

					if ts.TypeParams != nil {

						// for _, p := range ts.TypeParams.List {

						// 	for _, n := range p.Names {
						// 		t.TypeParameters =
						// 			append(
						// 				t.TypeParameters,
						// 				n.Name,
						// 			)
						// 	}
						// }
					}

					sp.Types =
						append(
							sp.Types,
							t,
						)
				}

				return false
			})
		}
	}

	return nil
}