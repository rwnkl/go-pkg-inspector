package ui

/*
import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
)

type MainWindow struct {
    app fyne.App
    win fyne.Window
}

func NewMainWindow() *MainWindow {

    a := app.New()

    w := a.NewWindow("Go Package Inspector")
    w.Resize(fyne.NewSize(1200, 800))

    tree := NewTree()
    details := NewDetails()

    split := container.NewHSplit(tree, details)
    split.Offset = 0.30

    toolbar := NewToolbar()
    status := NewStatusBar()

    content := container.NewBorder(
        toolbar,
        status,
        nil,
        nil,
        split,
    )

    w.SetMainMenu(BuildMenu())
    w.SetContent(content)

    return &MainWindow{
        app: a,
        win: w,
    }
}

func (m *MainWindow) Run() error {
    m.win.ShowAndRun()
    return nil
}
*/