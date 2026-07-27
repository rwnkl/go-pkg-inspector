package analyzer

import "go/ast"

func (w *Walker) TypeSpecs(fn func(*ast.TypeSpec)) {

	w.Files(func(file *ast.File) {

		ast.Inspect(file, func(n ast.Node) bool {

			ts, ok := n.(*ast.TypeSpec)

			if ok {
				fn(ts)
			}

			return true
		})
	})
}