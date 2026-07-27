package semantic

// type TypeKind int

// const (
// 	TypeStruct TypeKind = iota
// 	TypeInterface
// 	TypeAlias
// 	TypeDefined
// )

type Type struct {
    SymbolBase

    Kind TypeKind

    Fields []*Field
    Methods []*Method
}

// func (t *Type) Kind() SymbolKind {
//     return SymbolType
// }

// func (t *Type) Name() string {
//     return t.SymbolName
// }

// func (t *Type) Position() Position {
//     return t.Position
// }

// func (t *Type) Exported() bool {
//     return t.Exported
// }