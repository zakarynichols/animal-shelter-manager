package utils

import (
	"encoding/json"
	"net/http"
	"runtime/debug"
)

type AppJsonError struct {
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
}

func AppHttpError(w http.ResponseWriter, err AppJsonError, code int) {
	debug.PrintStack()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if (code == 0) {
		err.StatusCode = http.StatusInternalServerError
	} else {
		err.StatusCode = code
	}
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}