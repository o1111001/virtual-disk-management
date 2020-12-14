package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/o1111001/virtual-disk-management/server/disks"
)

type HTTPPortNumber int

type APIServer struct {
	Port HTTPPortNumber
	Handler disks.HTTPHandlerFunc
	server *http.Server
}

func (s *APIServer) Start() error {
	if s.Handler == nil {
		return fmt.Errorf(" HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/disks", s.Handler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *APIServer) Stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}
	return s.server.Shutdown(context.Background())
}
