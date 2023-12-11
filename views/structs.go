package views

// Response-struct for JSON-response
type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	Username string `json:"username"`
	Mail     string `json:"mail"`
}
