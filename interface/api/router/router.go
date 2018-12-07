package router

import (
	"blog-api/interface/api/rest"
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/gorilla/mux"
)

func Route(gh *handler.Handler, rh rest.RestHandler) http.Handler {
	router := mux.NewRouter()

	// Health Check
	router.Path("/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) },
	).Methods(http.MethodGet)

	router.Path("/graphql").Handler(gh)

	v1 := router.PathPrefix("/v1").Subrouter()
	v1.Path("/contents").HandlerFunc(rh.CreateContent).Methods(http.MethodPost)

	return router
}
