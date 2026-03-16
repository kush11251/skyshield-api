package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/threats", getThreats).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getThreats(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]string{"threat1", "threat2"})
}