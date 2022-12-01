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
	router.HandleFunc("PUT", "/...", handleMethodNotAllowed)
	router.HandleFunc("DELETE", "/...", handleMethodNotAllowed)

	// add a healthcheck
	router.HandleFunc("GET", "/healthcheck", handleHealthCheck)

	// events
	router.HandleFunc("GET", "/event", handleNotDoneYet)
	router.HandleFunc("POST", "/event/", handleNotDoneYet)
	router.HandleFunc("GET", "/event/:eventID", handleNotDoneYet)

	// venues
	router.HandleFunc("GET", "/venue", handleNotDoneYet)
	router.HandleFunc("POST", "/venue/", handleNotDoneYet)
	router.HandleFunc("GET", "/venue/:venueID", handleNotDoneYet)

	// bands
	router.HandleFunc("GET", "/band", handleNotDoneYet)
	router.HandleFunc("POST", "/band/", handleNotDoneYet)
	router.HandleFunc("GET", "/band/:bandID", handleNotDoneYet)

	// promoters
	router.HandleFunc("GET", "/promoter", handleNotDoneYet)
	router.HandleFunc("POST", "/promoter/", handleNotDoneYet)
	router.HandleFunc("GET", "/promoter/:promoterID", handleNotDoneYet)

	log.Fatalln(http.ListenAndServe(":8081", router))
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
