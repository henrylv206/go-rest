package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)


var server http.Server

func main() {

	log.Println("Server starting ...")

	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")

	server = http.Server{Addr: ":8000", Handler: router}

	log.Fatal(server.ListenAndServe())
}

// test
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{'version': 'v1'}")
}

// shutdown
func ShutdownHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("shutdown ...")

	server.Shutdown(nil)

	json.NewEncoder(w).Encode("{'status': 'success'}")
}