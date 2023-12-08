package main

import (
	"GOrepos/GoAPI/controller"
	"net/http"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
