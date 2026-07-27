package analyzer

// Collector is implemented by every semantic collector.
type Collector interface {
	Collect(*Context) error
}