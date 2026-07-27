package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MainWindow struct {
	app fyne.App

	window fyne.Window

	tree *TreeView

	details *DetailsView

	status *StatusBar
}

func NewMainWindow(app fyne.App) *MainWindow {

	mw := &MainWindow{
		app: app,
	}

	mw.window = app.NewWindow("Go Package Inspector")

	mw.tree = NewTreeView()

	mw.details = NewDetailsView()

	mw.status = NewStatusBar()

	split := container.NewHSplit(
		mw.tree.CanvasObject(),
		mw.details.CanvasObject(),
	)

	split.Offset = 0.30

	content := container.NewBorder(
		nil,
		mw.status.CanvasObject(),
		nil,
		nil,
		split,
	)

	mw.window.SetContent(content)

	mw.window.Resize(fyne.NewSize(1200, 800))

	mw.BuildMenu()

	return mw
}

func (m *MainWindow) Show() {
	m.window.Show()
}

func (m *MainWindow) Window() fyne.Window {
	return m.window
}

func (m *MainWindow) Tree() *TreeView {
	return m.tree
}

func (m *MainWindow) Details() *DetailsView {
	return m.details
}

func (m *MainWindow) StatusBar() *StatusBar {
	return m.status
}