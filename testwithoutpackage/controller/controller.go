package controller

import (
	"testwithoutpackage/service"

	"github.com/gorilla/mux"
)

func ApiRoutes(router *mux.Router) {
	router.Use(service.CommonMiddleware)
	router.HandleFunc("/", service.HomeLink)
	router.HandleFunc("/event", service.CreateEvent).Methods("POST")
	router.HandleFunc("/events", service.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", service.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", service.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", service.DeleteEvent).Methods("DELETE")

}
