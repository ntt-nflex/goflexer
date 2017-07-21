package goflexer

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Event holds the flexer event data
type Event struct {
	params map[string]interface{}
	raw    json.RawMessage
}

// NewEvent create a new Event object
func NewEvent(raw json.RawMessage) Event {
	return Event{
		raw: raw,
	}
}

// Unmarshal parses the event data and stores the result
// in the value pointed to by v.
func (e *Event) Unmarshal(v interface{}) error {
	return json.Unmarshal(e.raw, v)
}

// Get ...
func (e *Event) Get(key string) (interface{}, bool) {
	// unpack params only on first request to access
	if e.params == nil {
		err := json.Unmarshal(e.raw, &e.params)
		if err != nil {
			fmt.Printf("ERROR: %s", err)
			return nil, false
		}
	}

	value, ok := e.params[key]
	return value, ok
}

// GetString ...
func (e *Event) GetString(key string) (string, error) {

	value, ok := e.Get(key)
	if ok {
		retval, ok := value.(string)
		if ok {
			return retval, nil
		}
		return "", errors.New("Value is not string")
	}
	return "", errors.New("Key not found")
}

// GetInt ...
func (e *Event) GetInt(key string) (int64, error) {
	value, ok := e.Get(key)
	if ok {
		switch v := value.(type) {
		case int64:
			return v, nil
		case float64:
			return int64(v), nil
		default:
			return 0, errors.New("Value is not int64")
		}
	}
	return 0, errors.New("Key not found")
}

// GetFloat ...
func (e *Event) GetFloat(key string) (float64, error) {
	value, ok := e.Get(key)
	if ok {
		retval, ok := value.(float64)
		if ok {
			return retval, nil
		}
		return 0.0, errors.New("Value is not float64")
	}
	return 0.0, errors.New("Key not found")
}
