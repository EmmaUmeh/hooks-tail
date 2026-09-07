package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)



func HandleRoutes(r *chi.Mux) {
	// global middleware
	r.Use(chimiddleware.StripSlashes)

	
}