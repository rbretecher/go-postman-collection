package postman

type ListenType string

const (
	//prerequest script.
	Prerequest ListenType = "prerequest"
	//test script.
	Test ListenType = "test"
)

//acutal script object inside a postman collection
type Script struct {
	ScriptType string `json:"type"`
	Exec []string `json:"exec"`
}

// Event stores data that represents pre-quest and test scripts
type Event struct {
	Listen         ListenType   `json:"listen"`
	EventScript    *Script       `json:"script"`
}

func CreateEvent(listenType ListenType, script []string) *Event {
	return &Event{
		Listen:      listenType,
		EventScript: &Script{
			ScriptType: "text/javascript",
			Exec: script,
		},
	}
}
