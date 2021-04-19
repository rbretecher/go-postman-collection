package postman

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEventForPreRequestScript(t *testing.T) {
	scripts := []string{
		"",
	}

	//test the CreateEvent method
	testEvent := CreateEvent(EventType(EventTypePreRequestScript), "text/javascript", scripts)

	b, _ := json.Marshal(testEvent)
	testEventJSON := string(b)

	//expected outcome
	expectedEvent := &Event{
		Listen: EventType(EventTypePreRequestScript),
		Script: EventScript{
			ScriptType: "text/javascript",
			ScriptText: scripts,
		},
	}
	a, _ := json.Marshal(expectedEvent)
	expectedEventJSON := string(a)

	//compare expected outcome with CreateEvent response
	assert.Equal(t, expectedEvent, testEvent)
	assert.Equal(t, expectedEventJSON, testEventJSON)
}

func TestCreateEventForTests(t *testing.T) {
	scripts := []string{
		`
pm.test("Status code is 200", function () {
	pm.response.to.have.status(200);
});`,
		`
pm.test("Response time is less than 200ms", function () {
    pm.expect(pm.response.responseTime).to.be.below(200);
});`,
	}

	//test the CreateEvent method
	testEvent := CreateEvent(EventType(EventTypeTest), "text/javascript", scripts)

	b, _ := json.Marshal(testEvent)
	testEventJSON := string(b)

	//expected outcome
	expectedEvent := &Event{
		Listen: EventType(EventTypeTest),
		Script: EventScript{
			ScriptType: "text/javascript",
			ScriptText: scripts,
		},
	}
	a, _ := json.Marshal(expectedEvent)
	expectedEventJSON := string(a)

	//compare expected outcome with CreateEvent response
	assert.Equal(t, expectedEvent, testEvent)
	assert.Equal(t, expectedEventJSON, testEventJSON)
}
