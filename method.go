package postman

type Method int

const (
	GET      Method = iota
	PUT      Method = iota
	POST     Method = iota
	PATCH    Method = iota
	DELETE   Method = iota
	COPY     Method = iota
	HEAD     Method = iota
	OPTIONS  Method = iota
	LINK     Method = iota
	UNLINK   Method = iota
	PURGE    Method = iota
	LOCK     Method = iota
	UNLOCK   Method = iota
	PROPFIND Method = iota
	VIEW     Method = iota
)

func (m Method) String() string {
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
