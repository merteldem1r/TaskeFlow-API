package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Count  int         `json:"count,omitempty"`
	Error  string      `json:"error,omitempty"`
}

func JSON(w http.ResponseWriter, statusCode int, resp APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
