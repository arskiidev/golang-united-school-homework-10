package handlers

import (
	"net/http"
	"strconv"
)

func PostHeaders(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	a, _ := strconv.Atoi(headers["A"][0])
	b, _ := strconv.Atoi(headers["B"][0])
	w.Header().Add("a+b", strconv.Itoa(a+b))
}
