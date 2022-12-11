package main

import (
	"log"
	"net/http"
	"testwithoutpackage/controller"

	"github.com/gorilla/mux"
)

func main() {
	//initEvents()
	router := mux.NewRouter().StrictSlash(true)
	controller.ApiRoutes(router)
	log.Fatal(http.ListenAndServe(":8080", router))
}
