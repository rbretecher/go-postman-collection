package postman

import "testing"

func TestString(t *testing.T) {

	tables := []struct {
		method       method
		methodString string
	}{
		{Get, "GET"},
		{Put, "PUT"},
		{Post, "POST"},
		{Patch, "PATCH"},
		{Delete, "DELETE"},
		{Copy, "COPY"},
		{Head, "HEAD"},
		{Options, "OPTIONS"},
		{Link, "LINK"},
		{Unlink, "UNLINK"},
		{Purge, "PURGE"},
		{Lock, "LOCK"},
		{Unlock, "UNLOCK"},
		{Propfind, "PROPFIND"},
		{View, "VIEW"},
	}

	for _, table := range tables {
		if table.method.String() != table.methodString {
			t.Errorf("String was incorrect, got: %s, want: %s.", table.method.String(), table.methodString)
		}
	}
}
