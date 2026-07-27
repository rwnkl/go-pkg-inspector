package ui

import (
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TreeNode struct {
	ID       string
	ParentID string
	Text     string
	Leaf      bool
}

type TreeView struct {
	tree *widget.Tree

	nodes map[string]*TreeNode

	children map[string][]string

	onSelected func(id string)
}

func NewTreeView() *TreeView {

	tv := &TreeView{
		nodes:    make(map[string]*TreeNode),
		children: make(map[string][]string),
	}

	tv.tree = widget.NewTree(

		func(uid widget.TreeNodeID) []widget.TreeNodeID {

			children := tv.children[uid]

			result := make([]string, len(children))
			copy(result, children)

			sort.Strings(result)

			return result
		},

		func(uid widget.TreeNodeID) bool {

			n, ok := tv.nodes[uid]

			if !ok {
				return false
			}

			return !n.Leaf
		},

		func(branch bool) fyne.CanvasObject {

			return widget.NewLabel("")
		},

		func(uid widget.TreeNodeID, branch bool, obj fyne.CanvasObject) {

			label := obj.(*widget.Label)

			if n, ok := tv.nodes[uid]; ok {
				label.SetText(n.Text)
			} else {
				label.SetText(uid)
			}
		},
	)

	tv.tree.OnSelected = func(uid widget.TreeNodeID) {

		if tv.onSelected != nil {
			tv.onSelected(uid)
		}
	}

	return tv
}

func (t *TreeView) CanvasObject() fyne.CanvasObject {
	return t.tree
}

func (t *TreeView) SetOnSelected(fn func(id string)) {
	t.onSelected = fn
}

func (t *TreeView) Clear() {

	t.nodes = make(map[string]*TreeNode)
	t.children = make(map[string][]string)

	t.tree.Refresh()
}

func (t *TreeView) Add(node *TreeNode) {

	t.nodes[node.ID] = node

	t.children[node.ParentID] =
		append(
			t.children[node.ParentID],
			node.ID,
		)
}

func (t *TreeView) Refresh() {
	t.tree.Refresh()
}

func (t *TreeView) Open(id string) {
	t.tree.OpenBranch(id)
}

func (t *TreeView) Close(id string) {
	t.tree.CloseBranch(id)
}

func (t *TreeView) Select(id string) {
	t.tree.Select(id)
}


/*
Example
tree := ui.NewTreeView()

tree.Add(&ui.TreeNode{
	ID:   "",
	Text: "Workspace",
})

tree.Add(&ui.TreeNode{
	ID:       "pkg:fmt",
	ParentID: "",
	Text:     "fmt",
})

tree.Add(&ui.TreeNode{
	ID:       "type:Stringer",
	ParentID: "pkg:fmt",
	Text:     "Stringer",
	Leaf:     true,
})

tree.Refresh()

I would improve this before adding the symbol browser

While this implementation works, I wouldn't keep ID and ParentID as plain strings for long.

Earlier, we introduced the idea of stable semantic.IDs. I'd carry that through to the UI:

type TreeNode struct {
    ID       semantic.ID
    ParentID semantic.ID

    Text string
    Leaf bool
}

The TreeView would internally convert those IDs to widget.
TreeNodeID (which is currently a string in Fyne), while the rest of the application would work exclusively with strongly typed semantic IDs. 
That eliminates accidental ID mix-ups and makes the UI and analyzer integrate more cleanly.
*/