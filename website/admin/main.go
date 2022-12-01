package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PunkRockProgramming/diy-dot-org/internal/way"
)

func main() {
	router := way.NewRouter()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "There is nothing here")
	})
	// we won't need these Methods
	router.HandleFunc("POST", "/...", handleMethodNotAllowed)
	router.HandleFunc("PUT", "/...", handleMethodNotAllowed)
	router.HandleFunc("DELETE", "/...", handleMethodNotAllowed)

	// add a healthcheck
	router.HandleFunc("GET", "/healthcheck", handleHealthCheck)

	log.Fatalln(http.ListenAndServe(":8083", router))
}

func handleMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r.Context())
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Don't even try it.")
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r.Context())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Things are looking good.")
}

func handleNotDoneYet(w http.ResponseWriter, r *http.Request) {
	log.Printf("%+v", r.Context())
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "We're working on it.")
}
