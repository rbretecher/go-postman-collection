package postman

type method string

const (
	// Get HTTP method.
	Get method = "GET"
	// Put HTTP method.
	Put method = "PUT"
	// Post HTTP method.
	Post method = "POST"
	// Patch HTTP method.
	Patch method = "PATCH"
	// Delete HTTP method.
	Delete method = "DELETE"
	// Copy HTTP method.
	Copy method = "COPY"
	// Head HTTP method.
	Head method = "HEAD"
	// Options HTTP method.
	Options method = "OPTIONS"
	// Link HTTP method.
	Link method = "LINK"
	// Unlink HTTP method.
	Unlink method = "UNLINK"
	// Purge HTTP method.
	Purge method = "PURGE"
	// Lock HTTP method.
	Lock method = "LOCK"
	// Unlock HTTP method.
	Unlock method = "UNLOCK"
	// Propfind HTTP method.
	Propfind method = "PROPFIND"
	// View HTTP method.
	View method = "VIEW"
)
