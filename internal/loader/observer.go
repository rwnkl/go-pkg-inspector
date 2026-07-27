package loader

// EventKind identifies the type of loader event.
type EventKind int

const (
	EventMessage EventKind = iota
	EventModule
	EventPackage
	EventFile
	EventDiagnostic
)

// Event describes one loader event.
type Event struct {
	Kind EventKind

	// Human-readable message.
	Message string

	// Optional module/package/file name.
	Name string

	// Progress counters (0 if unknown).
	Current int
	Total   int

	// Optional diagnostic.
	Diagnostic *Diagnostic
}

// Observer receives loader events.
type Observer interface {
	OnEvent(Event)
}