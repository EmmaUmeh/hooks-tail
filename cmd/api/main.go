package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"hook-tail/internal/database"
	"hook-tail/internal/routes"
)

func main() {
	log.SetReportCaller(true)

	router := gin.Default()

	database.InitialiseDB()
	routes.SetupRoutes(router)

	fmt.Println("Starting server on port 8080...")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatal(err)
	}
}