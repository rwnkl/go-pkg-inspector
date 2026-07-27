package analyzer

import (
	"go/ast"

	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

type StructCollector struct{}

func (*StructCollector) Collect(
	ctx *Context,
) error {

	for _, pkg := range ctx.Model.Packages {

		for _, t := range pkg.Types {

			if t.Kind != semantic.TypeStruct {
				continue
			}
		}
	}

	for i, wp := range ctx.Workspace.Packages {

		sp := ctx.Model.Packages[i]

		for _, file := range wp.AST {

			ast.Inspect(file, func(n ast.Node) bool {

				ts, ok := n.(*ast.TypeSpec)

				if !ok {
					return true
				}

				st, ok := ts.Type.(*ast.StructType)

				if !ok {
					return false
				}

				current := sp.Type(ts.Name.Name)

				if current == nil {
					return false
				}

				for _, field := range st.Fields.List {

					sf := &semantic.Field{}

					if len(field.Names) == 0 {

						sf.Embedded = true

						switch e := field.Type.(type) {

						case *ast.Ident:

							sf.Name = e.Name
							// sf.Type = e.Name

						case *ast.SelectorExpr:

							// sf.Type = "selector"
						}

					} else {

						sf.Name = field.Names[0].Name

						// switch t := field.Type.(type) {

						// case *ast.Ident:

						// 	sf.Type = t.Name

						// default:

						// 	sf.Type = "<complex>"
						// }
					}

					if field.Tag != nil {

						sf.Tag = field.Tag.Value
					}

					current.Fields =
						append(
							current.Fields,
							sf,
						)
				}

				return false

			})
		}
	}

	return nil
}