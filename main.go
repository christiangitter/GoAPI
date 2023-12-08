package main

import(
	"net/http"
	"encoding/json"
) 

// Response-struct for JSON-response
type Response struct {
	Code int `json:"code"`
	Body interface{} `json:"body"`
}

//main-function
func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := Response{
				Code: http.StatusOK,
				Body: "healthy",
			}
			json.NewEncoder(w).Encode(data)
		}
	})

	http.ListenAndServe(":3000", mux)
}