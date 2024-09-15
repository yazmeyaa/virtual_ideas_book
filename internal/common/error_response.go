package common

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

const (
	UNREACHABLE_BODY_ERROR = "Invalid request payload"
)

func WriteError(w http.ResponseWriter, code int, err string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{
		Error: err,
	})
}
