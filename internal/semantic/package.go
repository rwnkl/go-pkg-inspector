package semantic

type Package struct {

	Name string

	Path string

	Types []*Type

	Functions []*Function

	Constants []*Constant

	index *Index
}
