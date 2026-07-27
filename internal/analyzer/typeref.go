package analyzer

import (
    "go/ast"

    "github.com/rwnkl/go-package-inspector/internal/semantic"
)

func BuildTypeRef(expr ast.Expr) *semantic.TypeRef {

    switch t := expr.(type) {

    case *ast.Ident:
        return &semantic.TypeRef{
            Kind: semantic.TypeNamed,
            Name: t.Name,
        }

    case *ast.StarExpr:
        return &semantic.TypeRef{
            Kind: semantic.TypePointer,
            Element: BuildTypeRef(t.X),
        }

    case *ast.ArrayType:

        return &semantic.TypeRef{
            Kind: semantic.TypeSlice,
            Element: BuildTypeRef(t.Elt),
        }

    case *ast.MapType:

        return &semantic.TypeRef{
            Kind: semantic.TypeMap,
            Key: BuildTypeRef(t.Key),
            Value: BuildTypeRef(t.Value),
        }

    default:

        return &semantic.TypeRef{
            Name: ExprString(nil, expr),
        }
    }
}