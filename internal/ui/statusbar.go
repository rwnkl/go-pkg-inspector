package ui

import (
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StatusBar is shown at the bottom of the main window.
type StatusBar struct {
	mu sync.Mutex

	root *fyne.Container

	label *widget.Label

	progress *widget.ProgressBarInfinite
}

// NewStatusBar creates a new status bar.
func NewStatusBar() *StatusBar {

	l := widget.NewLabel("Ready")

	p := widget.NewProgressBarInfinite()
	p.Hide()

	root := container.NewBorder(
		nil,
		nil,
		nil,
		p,
		l,
	)

	return &StatusBar{
		root:     root,
		label:    l,
		progress: p,
	}
}

// CanvasObject returns the widget.
func (s *StatusBar) CanvasObject() fyne.CanvasObject {
	return s.root
}

// Text returns the current message.
func (s *StatusBar) Text() string {

	s.mu.Lock()
	defer s.mu.Unlock()

	return s.label.Text
}

// SetText changes the status text.
func (s *StatusBar) SetText(text string) {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.label.SetText(text)
}

// BeginBusy displays the animated progress indicator.
func (s *StatusBar) BeginBusy() {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.progress.Show()
	s.progress.Start()
}

// EndBusy hides the progress indicator.
func (s *StatusBar) EndBusy() {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.progress.Stop()
	s.progress.Hide()
}

// Reset resets the status bar.
func (s *StatusBar) Reset() {

	s.mu.Lock()
	defer s.mu.Unlock()

	s.progress.Stop()
	s.progress.Hide()

	s.label.SetText("Ready")
}


/*
I'd define a single immutable status value:

type Status struct {
    Text string
    Busy bool
}

and update the widget through one method:

func (s *StatusBar) Set(status Status)

That gives you atomic updates (text and busy state change together), 
simplifies event handling from the controller, and makes future extensions (warnings, errors, progress percentage, elapsed time) 
much easier without changing the public API.
*/