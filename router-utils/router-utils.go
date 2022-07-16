package routerutils

import (
	"net/http"

	"github.com/GolangUnited/helloweb/rest"
	"github.com/gorilla/mux"
)

func Create() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/name/{PARAM}", rest.Get).Methods(http.MethodGet)
	router.HandleFunc("/bad", rest.BadGet).Methods(http.MethodGet)
	router.HandleFunc("/data", rest.PostData).Methods(http.MethodPost)
	router.HandleFunc("/headers", rest.PostHeader).Methods(http.MethodPost)

	return router
}
