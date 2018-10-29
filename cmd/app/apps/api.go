package apps

import (
	"net/http"
	"log"
	"encoding/json"
)

// api controller


func CreateAppsHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("create apps ...")



	json.NewEncoder(w).Encode("{'create': 'apps'}")

}