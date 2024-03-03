package api

import (
	"encoding/json"
	"net/http"
)

func Send(w http.ResponseWriter, r *http.Request, body interface{}, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(body)
}
