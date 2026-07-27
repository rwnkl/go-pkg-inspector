package semantic

// Model is the root semantic representation of a workspace.
type Model struct {
	Packages []*Package
}