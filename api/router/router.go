package router

import (
	"music-backend/api/handlers"

	"github.com/julienschmidt/httprouter"
)

func SetupRouter() *httprouter.Router{
	r := httprouter.New()

	r.GET("/", handlers.IndexHandler)

	return r
}