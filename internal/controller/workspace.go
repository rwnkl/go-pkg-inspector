package controller

import (
	"context"
	"sync"

	"github.com/rwnkl/go-package-inspector/internal/loader"
)

type WorkspaceController struct {
	mu sync.Mutex

	loader *loader.Loader

	cancel context.CancelFunc

	events chan Event
}

func NewWorkspaceController(
	l *loader.Loader,
) *WorkspaceController {

	return &WorkspaceController{
		loader: l,
		events: make(chan Event, 16),
	}
}

// Events returns a read-only event stream.
func (c *WorkspaceController) Events() <-chan Event {
	return c.events
}

// Cancel aborts the current load operation.
func (c *WorkspaceController) Cancel() {

	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cancel != nil {
		c.cancel()
		c.cancel = nil
	}

}

// Load starts loading asynchronously.
func (c *WorkspaceController) Load(parent context.Context) {

	c.Cancel()

	ctx, cancel := context.WithCancel(parent)

	c.mu.Lock()
	c.cancel = cancel
	c.mu.Unlock()

	go c.load(ctx)
}

func (c *WorkspaceController) load(
	ctx context.Context,
) {

	c.events <- Event{
		Type: EventStarted,
	}

	c.progress(
		"Loading workspace...",
		0,
		0,
	)

	c.progress(
		"Reading packages...",
		0,
		0,
	)

	// result, err := c.loader.Load(ctx)
	result, err := c.loader.Load(ctx, nil)

	if err != nil {

		c.events <- Event{
			Type: EventFailed,
			Err:  err,
		}

		return
	}

	c.events <- Event{
		Type:      EventFinished,
		Result:    result,
		Workspace: result.Workspace,
	}

	if result.Workspace != nil {

		c.progress(
			"Workspace loaded",
			result.Workspace.PackageCount(),
			result.Workspace.PackageCount(),
		)
	}
}


func (c *WorkspaceController) progress(
	msg string,
	current int,
	total int,
) {

	c.events <- Event{
		Type: EventProgress,
		Progress: &Progress{
			Message: msg,
			Current: current,
			Total: total,
		},
	}
}
