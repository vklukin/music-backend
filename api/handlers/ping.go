package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"music-backend/services/logging"
)



func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ping"))
	logging.Log("Ping")
}