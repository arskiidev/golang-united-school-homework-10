package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func PostData(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	msg := fmt.Sprintf("I got message:\n%v", string(body))
	w.Write([]byte(msg))
}
