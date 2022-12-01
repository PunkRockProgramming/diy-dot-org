package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/PunkRockProgramming/diy-dot-org/internal/templates"
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

	// events
	router.HandleFunc("GET", "/event", handleNotDoneYet)
	router.HandleFunc("GET", "/event/:eventID", handleNotDoneYet)
	router.HandleFunc("GET", "/event/submit", handleNotDoneYet)

	// venues
	router.HandleFunc("GET", "/venue", handleNotDoneYet)
	router.HandleFunc("GET", "/venue/", handleNotDoneYet)
	router.HandleFunc("GET", "/venue/submit", handleNotDoneYet)

	// bands
	router.HandleFunc("GET", "/band", handleNotDoneYet)
	router.HandleFunc("GET", "/band/", handleNotDoneYet)
	router.HandleFunc("GET", "/band/submit", handleNotDoneYet)

	// promoters
	router.HandleFunc("GET", "/promoter", handleNotDoneYet)
	router.HandleFunc("GET", "/promoter/", handleNotDoneYet)
	router.HandleFunc("GET", "/promoter/submit", handleNotDoneYet)

	log.Fatalln(http.ListenAndServe(":80", router))
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
	simpleMessageTemplate, err := template.ParseFiles("templates/simple-message.html")
	if err != nil {
		log.Printf("err: %s", err.Error())
	}
	data := templates.SimpleMessage{
		OrgName: "OKC DIY",
		Title:   "Not Done Yet",
		Message: "coming soon!",
	}

	simpleMessageTemplate.Execute(w, data)
}
