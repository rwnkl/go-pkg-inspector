package analyzer

import (
	"go/ast"
	"go/token"

	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

func BuildSignature(
	fset *token.FileSet,
	ft *ast.FuncType,
) semantic.Signature {

	var sig semantic.Signature

	if ft == nil {
		return sig
	}

	if ft.Params != nil {

		for i, field := range ft.Params.List {

			p := semantic.Parameter{
				// Type: BuildTypeRef(
				// 	fset,
				// 	field.Type,
				// ),
			}

			if len(field.Names) != 0 {
				p.Name = field.Names[0].Name
			}

			if ft.Params.NumFields()-1 == i {

				_, sig.Variadic =
					field.Type.(*ast.Ellipsis)
			}

			sig.Parameters =
				append(
					sig.Parameters,
					p,
				)
		}
	}

	if ft.Results != nil {

		for _, field := range ft.Results.List {

			r := semantic.Parameter{
				// Type: BuildTypeRef(
				// 	fset,
				// 	field.Type,
				// ),
			}

			if len(field.Names) != 0 {
				r.Name = field.Names[0].Name
			}

			sig.Results =
				append(
					sig.Results,
					r,
				)
		}
	}

	if ft.TypeParams != nil {

		for _, field := range ft.TypeParams.List {

			for _, n := range field.Names {

				sig.TypeParameters =
					append(
						sig.TypeParameters,
						n.Name,
					)
			}
		}
	}

	return sig
}