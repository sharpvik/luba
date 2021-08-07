package main

import (
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer() *Server {
	return &Server{
		&http.Server{
			Addr:         config.LubaAddr(),
			Handler:      NewRuntime().Handler(),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  10 * time.Second,
		},
	}
}
