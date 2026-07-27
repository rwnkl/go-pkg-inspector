package semantic

type Field struct {
    SymbolBase

    Name string

    // Human-readable representation
    TypeName string

    // Canonical compiler type
    // GoType types.Type

    Tag string

    Embedded bool
}