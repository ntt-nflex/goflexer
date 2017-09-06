package goflexer

// Handler represents a flexer handler
type Handler func(Event, *Context) (Result, error)

var handlers = map[string]Handler{}

// GetHandlers returns a map of registered handlers
func GetHandlers() map[string]Handler {
	return handlers
}

// RegisterHandler registers a function against a given handler name
func RegisterHandler(handler string, fn Handler) {
	handlers[handler] = fn
}
