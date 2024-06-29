package router

import (
	"music-backend/api/handlers"
	"music-backend/services/config"

	"github.com/julienschmidt/httprouter"
)

func routePath(path string) string{
	initialPath := config.Get().APIInitialPath

	return initialPath + path
}

func SetupRouter() *httprouter.Router{
	r := httprouter.New()

	r.GET("/", handlers.IndexHandler)
	r.GET(routePath("/ping"), handlers.Ping)

	return r
}