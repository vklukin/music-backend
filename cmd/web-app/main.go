package main

import (
	"fmt"
	"net"
	"net/http"

	"github.com/rs/cors"

	"music-backend/api/router"
	"music-backend/services/config"
	"music-backend/services/logging"
)

func main() {
	cfg := config.Get()

	r := router.SetupRouter()
	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowedMethods:      []string{"POST", "GET", "PUT", "OPTIONS"},
		AllowedHeaders:      []string{"*"},
		AllowCredentials:    false,
		AllowPrivateNetwork: true,
	})

	handler := c.Handler(r)
	srvAdress := fmt.Sprintf("0.0.0.0:%v", cfg.Port)
	server := &http.Server{
		Addr:    srvAdress,
		Handler: handler,
	}

	listener, err := net.Listen("tcp", srvAdress)
	if err != nil {
		logging.LogWithError(fmt.Sprintf("Could not listen on %d", cfg.Port), err.Error())
	}

	fmt.Printf("Server has been started on port %d!", cfg.Port)
	logging.Log(fmt.Sprintf("Server has been started on port %d!", cfg.Port))
	server.Serve(listener)
	// TODO: need to make gracefull shutdown
}
