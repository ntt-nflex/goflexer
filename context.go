package goflexer

// Context holds the flexer context
type Context struct {
	ModuleID    string
	Config      *Config
	Credentials interface{}
	Secrets     map[string]interface{}
	API         *CmpClient
	//State       *State
}

// NewContext creates a new context from a config
func NewContext(conf *Config) *Context {
	api := NewCmpClient(conf)

	c := Context{
		Config: conf,
		API:    api,
	}
	return &c
}

// TODO Add State Handler
