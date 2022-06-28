package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func PostHeaders(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	a, err := strconv.Atoi(headers["A"][0])
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.Atoi(headers["B"][0])
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Add("a+b", strconv.Itoa(a+b))
}
