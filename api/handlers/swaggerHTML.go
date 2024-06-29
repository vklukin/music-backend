package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"music-backend/services/config"
	"music-backend/services/logging"
)

func SwaggerHTML(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	cfg := config.Get()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, cfg.SwaggerHTMLFilePath)

	logging.Log("Render swagger UI. Endpoint: /swagger")
}