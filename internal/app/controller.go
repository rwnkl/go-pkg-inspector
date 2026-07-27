package app

import (
	"github.com/rwnkl/go-package-inspector/internal/semantic"
)

// This keeps the UI completely independent of the loader and analyzer implementation.
type Controller interface {

    OpenWorkspace(dir string) error

    Symbols() []semantic.Symbol

    // Events() <-chan Event
}