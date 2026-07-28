package semantic

type SymbolBase struct {
    ID ID

    Name string

    Exported bool

    Position Position
}

// func MakeID(kind, pkgPath, name string) semantic.ID {
//     return semantic.ID(kind + ":" + pkgPath + "." + name)
// }

func MakeID(kind, pkgPath, name string) ID {
    return ID(kind + ":" + pkgPath + "." + name)
}
