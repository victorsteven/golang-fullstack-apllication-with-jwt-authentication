package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := userRoutes
	return routes
}

func SetUpRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}
