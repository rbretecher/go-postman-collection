package postman

type method int

const (
	// Get HTTP method.
	Get method = iota
	// Put HTTP method.
	Put method = iota
	// Post HTTP method.
	Post method = iota
	// Patch HTTP method.
	Patch method = iota
	// Delete HTTP method.
	Delete method = iota
	// Copy HTTP method.
	Copy method = iota
	// Head HTTP method.
	Head method = iota
	// Options HTTP method.
	Options method = iota
	// Link HTTP method.
	Link method = iota
	// Unlink HTTP method.
	Unlink method = iota
	// Purge HTTP method.
	Purge method = iota
	// Lock HTTP method.
	Lock method = iota
	// Unlock HTTP method.
	Unlock method = iota
	// Propfind HTTP method.
	Propfind method = iota
	// View HTTP method.
	View method = iota
)

func (m method) String() string {
	names := []string{
		"GET",
		"PUT",
		"POST",
		"PATCH",
		"DELETE",
		"COPY",
		"HEAD",
		"OPTIONS",
		"LINK",
		"UNLINK",
		"PURGE",
		"LOCK",
		"UNLOCK",
		"PROPFIND",
		"VIEW",
	}

	return names[m]
}
