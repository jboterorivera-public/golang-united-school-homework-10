package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]

	w.Write([]byte("Hello, " + param + "!"))
	w.WriteHeader(http.StatusOK)
}

func BadGet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}
