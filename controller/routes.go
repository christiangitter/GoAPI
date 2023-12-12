package controller

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheck())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/readAll", readAll())
	return mux
}
