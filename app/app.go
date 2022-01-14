package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/data", getData)

	log.Fatal(http.ListenAndServe(":3333", router))
}
