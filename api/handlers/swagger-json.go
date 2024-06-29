package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"music-backend/services/config"
	"music-backend/services/logging"
)

func SwaggerJson(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cfg := config.Get()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, cfg.SwaggerFilePath)
	
	logging.Log("Send swagger.json file. Endpoint: /swagger.json")
}