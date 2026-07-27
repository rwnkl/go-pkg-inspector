package app

import (
    fyneapp "fyne.io/fyne/v2/app"

    "github.com/rwnkl/go-package-inspector/internal/ui"
)


// Notice that Run() returns an error even though Fyne currently doesn't produce one. 
// This makes it easier to add configuration loading or startup validation later.
func Run() error {

    a := fyneapp.New()

    window := ui.NewMainWindow(a)

    window.Show()

    a.Run()

    return nil
}
