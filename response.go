package postman

// A Response represents an HTTP request and response.
type Response struct {
	Name            string    `json:"name"`
	OriginalRequest *Request  `json:"originalRequest"`
	Status          string    `json:"status,omitempty"`
	Code            int       `json:"code,omitempty"`
	PreviewLanguage string    `json:"_postman_previewlanguage,omitempty"`
	Header          []*Header `json:"header,omitempty"`
	Body            string    `json:"body,omitempty"`
}
