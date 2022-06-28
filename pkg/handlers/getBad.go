package handlers

import "net/http"

func GetBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}
