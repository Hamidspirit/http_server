package main

import "net/http"

func handleError(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, http.StatusBadRequest, "Something went wrong")
}
