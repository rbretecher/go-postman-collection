package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTest(t *testing.T) {
	assert.Equal(
		t,
		&Event{
			Listen: test,
			EventScript: &Script{
				ScriptType: "text/javascript",
				Exec:[]string{"console.log(\"foo\")"},
			},
		},
		CreateEvent(test,[]string{"console.log(\"foo\")"}),
	)
}
