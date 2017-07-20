package main

import (
	"flag"
	"plugin"
	"reflect"

	log "github.com/Sirupsen/logrus"
	"github.com/ntt-nflex/goflexer"
)

func main() {
	var handler, module string
	flag.StringVar(&handler, "handler", "Test", "Handler method to run")
	flag.StringVar(&module, "module", "plugin.so", "Plugin module to load")
	flag.Parse()

	conf := goflexer.NewConfigFromYAML()
	log.Infof("Loaded config: %+v", conf)

	context := goflexer.NewContext(conf)
	event := goflexer.Event{}

	log.Infof("Importing go plugin module \"%s\"", module)
	p, err := plugin.Open(module)
	if err != nil {
		log.Fatalf("Could not import go plugin module %s: %s", module, err)
	}

	s, err := p.Lookup(handler)
	if err != nil {
		log.Fatalf("Method not found: %s - %s", handler, err)
	}

	// Ensure the handler is a func with correct interface
	f, ok := s.(func(goflexer.Context, goflexer.Event) goflexer.Result)
	if !ok {
		log.Fatalf("Incorrect method definition: %s %s", handler, reflect.TypeOf(s))
	}

	// Run it
	flexResult := f(context, event)

	log.Infof("Module Execution Completed: %+v", flexResult)
}
