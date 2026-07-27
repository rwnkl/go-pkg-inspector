package controller

import (
	"github.com/rwnkl/go-package-inspector/internal/loader"
	"github.com/rwnkl/go-package-inspector/internal/workspace"
)

type EventType int

const (
	EventStarted EventType = iota

	// emitted while loading
	EventProgress

	EventFinished

	EventFailed
)

type Progress struct {

	// Human readable message.
	Message string

	// Current item.
	Current int

	// Total items.
	Total int
}

func (p Progress) Percent() float64 {

	if p.Total == 0 {
		return 0
	}

	return float64(p.Current) / float64(p.Total)
}

type Event struct {

	Type EventType

	Workspace *workspace.Workspace

	Result *loader.Result

	Progress *Progress

	Err error
}