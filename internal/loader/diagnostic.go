package loader

import "fmt"

// Severity indicates the importance of a diagnostic.
type Severity int

const (
	SeverityInfo Severity = iota
	SeverityWarning
	SeverityError
)

func (s Severity) String() string {
	switch s {
	case SeverityInfo:
		return "info"
	case SeverityWarning:
		return "warning"
	case SeverityError:
		return "error"
	default:
		return "unknown"
	}
}

// Diagnostic represents one compiler or loader message.
type Diagnostic struct {
	Severity Severity

	Package string

	File string

	Line int

	Column int

	Message string
}

func (d Diagnostic) Error() string {

	if d.File == "" {
		return d.Message
	}

	return fmt.Sprintf(
		"%s:%d:%d: %s",
		d.File,
		d.Line,
		d.Column,
		d.Message,
	)
}