package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func StartServer() {

	r := mux.NewRouter()
	r.HandleFunc("/file/{fileName}", GetFiles)
	log.Fatal(http.ListenAndServe(":4000", r))

}
