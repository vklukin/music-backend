package handlers

import (
	"music-backend/services/config"
	"music-backend/services/logging"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SwaggerJson(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	cfg := config.Get()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	http.ServeFile(w, r, cfg.SwaggerFilePath)
	
	logging.Log("Send swagger.json file")
}