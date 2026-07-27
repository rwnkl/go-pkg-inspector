package analyzer

import "go/ast"

func receiverName(
	fd *ast.FuncDecl,
) string {

	if fd.Recv == nil {
		return ""
	}

	field := fd.Recv.List[0]

	switch t := field.Type.(type) {

	case *ast.Ident:
		return t.Name

	case *ast.StarExpr:

		if id, ok := t.X.(*ast.Ident); ok {
			return id.Name
		}
	}

	return ""
}