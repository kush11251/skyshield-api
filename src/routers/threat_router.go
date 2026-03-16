package routers

import (
	"github.com/gorilla/mux"

	"skyshield-api/src/controllers"
)

func ThreatRouter(r *mux.Router) {
	r.HandleFunc("/threats", controllers.GetThreats).Methods("GET")
}