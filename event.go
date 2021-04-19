package postman

// EventType is an enum type for an Event's `Listen` field.
type EventType string

const (
	// EventTypePreRequestScript is the enum value for a Pre-Request Script that will run against all `Item` in a `Collection`
	EventTypePreRequestScript string = "prerequest"
	// EventTypeTest is the enum value for Tests that will run against all `Item` in a `Collection`
	EventTypeTest string = "test"
)

// EventScript is a script (or group of scripts) that are run against a given `Item`.
type EventScript struct {
	ScriptType string   `json:"type"`
	ScriptText []string `json:"exec"`
}

// A Event is a Pre-request Script or a Test that can be run post-request.
type Event struct {
	Listen EventType   `json:"listen"`
	Script EventScript `json:"script"`
}

// CreateEvent creates a new Variable of type string.
// "text/javascript"
func CreateEvent(eventType EventType, scriptType string, scripts []string) *Event {
	return &Event{
		Listen: eventType,
		Script: EventScript{
			ScriptType: scriptType,
			ScriptText: scripts,
		},
	}
}
