package analyzer

import (
	"go/ast"

	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

type MethodCollector struct{}

func (*MethodCollector) Collect(ctx *Context) error {

	wp := ctx.Package.WorkspacePackage

	sp := ctx.Package.SemanticPackage

	walker := NewWalker(ctx)

	walker.FuncDecls(func(fd *ast.FuncDecl) {

		if fd.Recv == nil {
			return
		}

		receiver := receiverName(fd)

		t := sp.Type(receiver)

		if t == nil {
			return
		}

		m := &semantic.Method{
				Name: fd.Name.Name,
			}

		m.Exported = ast.IsExported(fd.Name.Name)

		m.Position =
				Position(
					wp.FileSet,
					fd.Pos(),
				)

		m.Receiver = receiver

		m.Signature =
			BuildSignature(
				wp.FileSet,
				fd.Type,
			)

		t.Methods =
			append(
				t.Methods,
				m,
			)

	})

	// for _, file := range wp.AST {

	// 	for _, decl := range file.Decls {

	// 		fd, ok := decl.(*ast.FuncDecl)

	// 		if !ok {
	// 			continue
	// 		}

	// 		if fd.Recv == nil {
	// 			continue
	// 		}

	// 		receiver := receiverName(fd)

	// 		t := sp.Type(receiver)

	// 		if t == nil {
	// 			continue
	// 		}

	// 		m := &semantic.Method{
	// 			Name: fd.Name.Name,
	// 		}

	// 		m.Exported = ast.IsExported(fd.Name.Name)

	// 		m.Position =
	// 			Position(
	// 				wp.FileSet,
	// 				fd.Pos(),
	// 			)

	// 		m.Receiver = receiver

	// 		// collectParameters(
	// 		// 	wp.FileSet,
	// 		// 	fd.Type.Params,
	// 		// 	&m.Parameters,
	// 		// )

	// 		// collectParameters(
	// 		// 	wp.FileSet,
	// 		// 	fd.Type.Results,
	// 		// 	&m.Results,
	// 		// )

	// 		m.Signature =
	// 			BuildSignature(
	// 				wp.FileSet,
	// 				fd.Type,
	// 		)

	// 		t.Methods =
	// 			append(
	// 				t.Methods,
	// 				m,
	// 			)
	// 	}
	// }

	return nil
}