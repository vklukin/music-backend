package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte(fmt.Sprintln("I glad to see you on this backend server")))
}