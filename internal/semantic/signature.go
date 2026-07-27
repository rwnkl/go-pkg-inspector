package semantic

// Signature describes the callable interface of a function or method.
type Signature struct {
	Parameters []Parameter

	Results []Parameter

	Variadic bool

	TypeParameters []string
}