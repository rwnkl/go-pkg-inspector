package ui

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/widget"
)

func NewTree() *widget.Tree {

    data := map[string][]string{
        "": {"workspace"},
        "workspace": {},
    }

    return widget.NewTree(

        func(uid widget.TreeNodeID) []widget.TreeNodeID {
            return data[uid]
        },

        func(uid widget.TreeNodeID) bool {
            _, ok := data[uid]
            return ok
        },

        func(branch bool) fyne.CanvasObject {
            return widget.NewLabel("")
        },

        func(uid widget.TreeNodeID, branch bool, obj fyne.CanvasObject) {
            obj.(*widget.Label).SetText(uid)
        },
    )
}
