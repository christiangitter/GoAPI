package views

// Response-struct for JSON-response
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}
