package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
)

func HttpServerError (res http.ResponseWriter, message string, code int) {
	
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(code)
	json.NewEncoder(res).Encode()
}