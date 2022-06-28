package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Write([]byte(fmt.Sprintf("Hello, %v!", vars["PARAM"])))
}
