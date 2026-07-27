package ui

import (
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Property represents one name/value pair displayed in the details view.
type Property struct {
	Name  string
	Value string
}

// DetailsView displays a list of properties.
type DetailsView struct {
	title *widget.Label

	grid *fyne.Container

	root *fyne.Container
}

// NewDetailsView creates a new details view.
func NewDetailsView() *DetailsView {

	title := widget.NewLabel("No Selection")

	grid := container.NewGridWithColumns(2)

	root := container.NewBorder(
		title,
		nil,
		nil,
		nil,
		container.NewScroll(grid),
	)

	return &DetailsView{
		title: title,
		grid:  grid,
		root:  root,
	}
}

// CanvasObject returns the root widget.
func (d *DetailsView) CanvasObject() fyne.CanvasObject {
	return d.root
}

// Clear removes all displayed properties.
func (d *DetailsView) Clear() {

	d.title.SetText("No Selection")

	d.grid.RemoveAll()
}

// Set displays a title and a list of properties.
func (d *DetailsView) Set(title string, properties []Property) {

	d.title.SetText(title)

	d.grid.RemoveAll()

	sort.Slice(properties, func(i, j int) bool {
		return properties[i].Name < properties[j].Name
	})

	for _, p := range properties {

		name := widget.NewLabel(p.Name + ":")

		value := widget.NewLabel(p.Value)
		value.Wrapping = fyne.TextWrapWord

		d.grid.Add(name)
		d.grid.Add(value)
	}
}


/*
Example
details := ui.NewDetailsView()

details.Set(
	"Type Person",
	[]ui.Property{
		{Name: "Package", Value: "github.com/example/demo"},
		{Name: "Kind", Value: "Struct"},
		{Name: "Exported", Value: "true"},
	},
)

The view will display:

Type Person

Exported:  true
Kind:       Struct
Package:    github.com/example/demo



Next improvement

I would not continue using a generic property list for long.

The inspector is going to show many different symbol types:

package
file
struct
interface
alias
enum
function
method
variable

Each has a different layout.

Instead, I'd introduce a renderer interface:

type DetailRenderer interface {
    CanRender(symbol semantic.Symbol) bool

    Render(symbol semantic.Symbol) fyne.CanvasObject
}

Then we'd implement renderers such as:

StructRenderer
InterfaceRenderer
FunctionRenderer
PackageRenderer
ConstantRenderer

The DetailsView would simply choose the appropriate renderer based on the selected symbol. This keeps each renderer focused on one symbol type and makes it easy to add richer views later (tables for fields, signatures, documentation, source locations, etc.) without turning DetailsView into a large switch statement.

At that point, the next major milestone would be wiring together:

MainWindow
TreeView
DetailsView
StatusBar
WorkspaceController

to produce the first runnable Go Package Inspector window that can load a workspace and respond to user selections.
*/