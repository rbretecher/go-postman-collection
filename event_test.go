package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEventCreateEvent(t *testing.T) {
	assert.Equal(
		t,
		&Event{
			Listen: Test,
			Script: &Script{
				Type: "text/javascript",
				Exec: []string{"console.log(\"foo\")"},
			},
		},
		CreateEvent(Test, []string{"console.log(\"foo\")"}),
	)
}
