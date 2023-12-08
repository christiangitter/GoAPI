package controller

import (
	"GOrepos/GoAPI/views"
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
			json.NewEncoder(w).Encode(data)
		}
	}
}
