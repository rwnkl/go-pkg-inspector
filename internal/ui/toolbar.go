package ui

import (
	// "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
    "fyne.io/fyne/v2/theme"
)

func NewToolbar() *widget.Toolbar {

    return widget.NewToolbar(
        widget.NewToolbarAction(
            theme.FolderOpenIcon(),
            func() {},
        ),
    )
}