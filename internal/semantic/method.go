package semantic

type Method struct {

	Name string

	Receiver string

	Signature Signature

	Exported bool

	Position Position
}