package semantic

// TypeKind describes the shape of a Go type.
type TypeKind int

const (
    TypeNamed TypeKind = iota
    TypePointer
    TypeSlice
    TypeArray
    TypeMap
    TypeChan
    TypeFunc
    TypeInterface
    TypeStruct
)

// TypeRef represents a Go type.
type TypeRef struct {
    Kind TypeKind

    // Named type
    Name string

    // Pointer/slice/array/channel element
    Element *TypeRef

    // Map key/value
    Key   *TypeRef
    Value *TypeRef
}