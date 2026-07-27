package model

type NodeKind int

const (
    RootNode NodeKind = iota
    PackageNode
    FileNode
    StructNode
    InterfaceNode
    FunctionNode
)

type Node struct {
    ID       string
    Name     string
    Kind     NodeKind
    Children []*Node
}