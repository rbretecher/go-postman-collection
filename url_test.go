package postman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLString(t *testing.T) {
	u := URL{
		Raw: "a-raw-url",
	}

	assert.Equal(t, "a-raw-url", u.String())
}
