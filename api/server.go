package api

import (
	"fmt"
	"golang_api_fullstack/api/router"
	"golang_api_fullstack/config"
	"log"
	"net/http"
)

func Run() {
	config.Load()
	fmt.Printf("running... at port %d", config.PORT)
	listen(config.PORT)

}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
