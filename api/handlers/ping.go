package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)



func Ping(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ping"))
}