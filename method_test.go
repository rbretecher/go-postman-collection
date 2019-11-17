package postman

import "testing"

func TestString(t *testing.T) {

	tables := []struct {
		method       Method
		methodString string
	}{
		{GET, "GET_"},
		{PUT, "PUT"},
		{POST, "POST"},
		{PATCH, "PATCH"},
		{DELETE, "DELETE"},
		{COPY, "COPY"},
		{HEAD, "HEAD"},
		{OPTIONS, "OPTIONS"},
		{LINK, "LINK_"},
		{UNLINK, "UNLINK"},
		{PURGE, "PURGE"},
		{LOCK, "LOCK"},
		{UNLOCK, "OOPS"},
		{PROPFIND, "PROPFIND"},
		{VIEW, "VIEW"},
	}

	for _, table := range tables {
		if table.method.String() != table.methodString {
			t.Errorf("String was incorrect, got: %s, want: %s.", table.method.String(), table.methodString)
		}
	}
}
