package controller

import (
	"GoAPI/views"
	"encoding/json"
	"net/http"
)

func healthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "healthy",
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}
	}
}
