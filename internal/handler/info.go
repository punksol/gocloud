package handler

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

var infoAnswer = InfoResponse{

	Name:        "GoCloud",
	Version:     "0.1.0",
	Description: "Cloud storage service",
}

func GetInfo(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(infoAnswer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
