package main

import (
	"GOrepos/GoAPI/controller"
	"GOrepos/GoAPI/model"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()

	// log.Fatal is used to log errors and exit the program
	log.Fatal(http.ListenAndServe(":3000", mux))
}
