package analyzer

import "go/ast"

func (w *Walker) GenDecls(fn func(*ast.GenDecl)) {

	w.Files(func(file *ast.File) {

		for _, decl := range file.Decls {

			gd, ok := decl.(*ast.GenDecl)

			if ok {
				fn(gd)
			}
		}
	})
}