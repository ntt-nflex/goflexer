package goflexer

type Context struct {
	Credentials interface{}
	Secrets     map[string]interface{}
}

type Event struct {
	params map[string]interface{}
}

type Result interface{}

func (e *Event) Get(key string) interface{} {
	return e.params[key]
}
