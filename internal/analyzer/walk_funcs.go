package analyzer

import "go/ast"

func (w *Walker) FuncDecls(fn func(*ast.FuncDecl)) {

	w.Files(func(file *ast.File) {

		for _, decl := range file.Decls {

			fd, ok := decl.(*ast.FuncDecl)

			if ok {
				fn(fd)
			}
		}
	})
}