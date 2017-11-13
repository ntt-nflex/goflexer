package main

import (
	"encoding/json"
	"flag"
	"plugin"
	"reflect"

	"github.com/ntt-nflex/goflexer"
	log "github.com/sirupsen/logrus"
)

func main() {
	var handler, module, testEvent string
	flag.StringVar(&handler, "handler", "Test", "Handler method to run")
	flag.StringVar(&module, "module", "plugin.so", "Plugin module to load")
	flag.StringVar(&testEvent, "event", "{}", "Test event (json string)")
	flag.Parse()

	conf := goflexer.NewConfigFromYAML()
	log.Infof("Loaded config: %+v", conf)

	context := goflexer.NewContext(conf)
	event := goflexer.NewEvent(json.RawMessage(testEvent))

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
	f, ok := s.(func(goflexer.Event, *goflexer.Context) goflexer.Result)
	if !ok {
		log.Fatalf("Incorrect method definition: %s %s", handler, reflect.TypeOf(s))
	}

	// Run it
	flexResult := f(event, context)

	log.Infof("Module Execution Completed: %+v", flexResult)
}
