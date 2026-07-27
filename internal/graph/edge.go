package graph

import "github.com/rwnkl/go-package-inspector/internal/semantic"

type EdgeKind int

const (
    EdgeContains EdgeKind = iota
    EdgeDefines
    EdgeImplements
    EdgeEmbeds
    EdgeCalls
    EdgeReferences
)

type Edge struct {
    From semantic.ID
    To   semantic.ID

    Kind EdgeKind
}