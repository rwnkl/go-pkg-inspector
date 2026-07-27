package loader

// ObserverFunc adapts a function to an Observer.
type ObserverFunc func(Event)

func (f ObserverFunc) OnEvent(e Event) {
	if f != nil {
		f(e)
	}
}