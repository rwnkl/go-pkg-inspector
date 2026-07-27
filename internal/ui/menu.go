package ui

import "fyne.io/fyne/v2"

func (m *MainWindow) BuildMenu() {

	fileMenu := fyne.NewMenu(
		"File",

		fyne.NewMenuItem(
			"Open Workspace...",
			func() {
				m.onOpenWorkspace()
			},
		),

		fyne.NewMenuItemSeparator(),

		fyne.NewMenuItem(
			"Exit",
			func() {
				m.app.Quit()
			},
		),
	)

	m.window.SetMainMenu(
		fyne.NewMainMenu(
			fileMenu,
		),
	)
}

func (m *MainWindow) onOpenWorkspace() {
	// implemented in next milestone
}