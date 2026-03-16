package controllers

import (
	"encoding/json"
	"net/http"

	"skyshield-api/src/models"
)

func GetThreats(w http.ResponseWriter, r *http.Request) {
	threats := []models.Threat{{ID: "1", Name: "threat1", Severity: "high"}, {ID: "2", Name: "threat2", Severity: "low"}}
	json.NewEncoder(w).Encode(threats)
}