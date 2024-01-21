package main

import (
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{
		"status": "ok",
	})
}

func handlerErr(w http.ResponseWriter, req *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal server error")
}
