package goflexer

// Event holds the flexer event data
type Event struct {
	params map[string]interface{}
}

// Result holds the flexer result
type Result interface{}

func (e *Event) Get(key string) interface{} {
	return e.params[key]
}
