package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
	"./app/apps"
)


var server http.Server

func main() {

	log.Println("Server starting ...")

	router := mux.NewRouter()

	// test & shutdown
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/shutdown", ShutdownHandler).Methods("GET")

	// TODO user group role permission app resource

	// app
	router.HandleFunc("/v1/apps", apps.CreateAppsHandler).Methods("GET")


	server = http.Server{Addr: ":8080", Handler: router}

	log.Fatal(server.ListenAndServe())
}

// test
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("path: / ...")

	json.NewEncoder(w).Encode("{'version': 'v1'}")
}

// shutdown
func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shutdown ...")

	server.Shutdown(nil)

	json.NewEncoder(w).Encode("{'status': 'success'}")
}git