package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateVariable(t *testing.T) {
	assert.Equal(
		t,
		&Variable{
			Name:  "a-name",
			Value: "a-value",
			Type:  "string",
		},
		CreateVariable("a-name", "a-value"),
	)
}
