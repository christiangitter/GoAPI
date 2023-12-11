package main

import (
	"GoAPI/controller"
	"GoAPI/model"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving on port 3000")

	// log.Fatal is used to log errors and exit the program
	log.Fatal(http.ListenAndServe(":3000", mux))
}
