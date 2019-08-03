package router

import (
	"golang_api_fullstack/api/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetUpRoutes(r)
}
