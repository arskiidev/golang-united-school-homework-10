package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
	}
	msg := fmt.Sprintf("I got message:\n%v", string(body))
	w.Write([]byte(msg))
}
