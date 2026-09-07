package main

import (
	"fmt"
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/go-chi/chi/v5"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handler.HandleRoutes(r)

	fmt.Println("Starting server on port 8080...")

	err := http.ListenAndServe(":8080", r)
	
	if err != nil {
		log.Fatal(err)
	}

}