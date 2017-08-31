package goflexer

type Handler func(Event, *Context) Result

var handlers = map[string]Handler{}

func GetHandlers() map[string]Handler {
	return handlers
}

func RegisterHandler(handler string, fn Handler) {
	handlers[handler] = fn
}
