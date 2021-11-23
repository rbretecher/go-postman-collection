package postman

// A Response represents an HTTP response.
type Response struct {
	ID              string      `json:"id,omitempty"`
	OriginalRequest *Request    `json:"originalRequest,omitempty"`
	ResponseTime    interface{} `json:"responseTime,omitempty"`
	Timings         interface{} `json:"timings,omitempty"`
	Headers         *HeaderList `json:"header,omitempty"`
	Cookies         []*Cookie   `json:"cookie,omitempty"`
	Body            string      `json:"body,omitempty"`
	Status          string      `json:"status,omitempty"`
	Code            int         `json:"code,omitempty"`
	Name            string      `json:"name,omitempty"`
	PreviewLanguage string      `json:"_postman_previewlanguage,omitempty"`
}
