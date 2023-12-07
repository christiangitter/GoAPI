package main

import (
	"fmt"
	"net/http"
	"Users/U782683/repos/GO/GoAPI/structs"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "healthy",
			}
			json.NewEncoder(w).Encode(data)
		}
	})

	http.ListenAndServe(":3000", mux)
}