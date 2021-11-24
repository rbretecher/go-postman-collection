package postman

// ListenType defines the kind of script attached to an event.
type ListenType string

const (
	// PreRequest script is usually executed before the HTTP request is sent.
	PreRequest ListenType = "prerequest"
	// Test script is usually executed after the actual HTTP request is sent, and the response is received.
	Test ListenType = "test"
)

// A Script is a snippet of Javascript code that can be used to to perform setup or teardown operations on a particular response.
type Script struct {
	ID   string   `json:"id,omitempty"`
	Type string   `json:"type,omitempty"`
	Exec []string `json:"exec,omitempty"`
	Src  *URL     `json:"src,omitempty"`
	Name string   `json:"name,omitempty"`
}

// An Event defines a script associated with an associated event name.
type Event struct {
	ID       string     `json:"id,omitempty"`
	Listen   ListenType `json:"listen,omitempty"`
	Script   *Script    `json:"script,omitempty"`
	Disabled bool       `json:"disabled,omitempty"`
}

// CreateEvent creates a new Event of type text/javascript.
func CreateEvent(listenType ListenType, script []string) *Event {
	return &Event{
		Listen: listenType,
		Script: &Script{
			Type: "text/javascript",
			Exec: script,
		},
	}
}
