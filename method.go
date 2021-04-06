package postman

type Method string

const (
	// Get HTTP Method.
	Get Method = "GET"
	// Put HTTP Method.
	Put Method = "PUT"
	// Post HTTP Method.
	Post Method = "POST"
	// Patch HTTP Method.
	Patch Method = "PATCH"
	// Delete HTTP Method.
	Delete Method = "DELETE"
	// Copy HTTP Method.
	Copy Method = "COPY"
	// Head HTTP Method.
	Head Method = "HEAD"
	// Options HTTP Method.
	Options Method = "OPTIONS"
	// Link HTTP Method.
	Link Method = "LINK"
	// Unlink HTTP Method.
	Unlink Method = "UNLINK"
	// Purge HTTP Method.
	Purge Method = "PURGE"
	// Lock HTTP Method.
	Lock Method = "LOCK"
	// Unlock HTTP Method.
	Unlock Method = "UNLOCK"
	// Propfind HTTP Method.
	Propfind Method = "PROPFIND"
	// View HTTP Method.
	View Method = "VIEW"
)
