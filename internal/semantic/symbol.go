package semantic

type SymbolKind int

const (
    SymbolPackage SymbolKind = iota
    SymbolType
    SymbolFunction
    SymbolMethod
    SymbolConstant
    SymbolVariable
)

type Symbol interface {

    ID() ID

    Kind() SymbolKind

    Name() string

    PackagePath() string

    Position() Position

    Exported() bool
}