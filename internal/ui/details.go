package ui

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func NewDetails() *fyne.Container {

    title := widget.NewLabelWithStyle(
        "Details",
        fyne.TextAlignLeading,
        fyne.TextStyle{Bold: true},
    )

    text := widget.NewMultiLineEntry()
    text.Disable()

    return container.NewBorder(
        title,
        nil,
        nil,
        nil,
        text,
    )
}