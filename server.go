package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ab22/abcd/config"
)

type Server struct {
	port         int
	router       *Router
	bootstrapper bootstrapper
}

func NewServer() *Server {
	s := &Server{}

	s.bootstrapper = bootstrapper{
		initializeConfigurationModule,
		initializeServicesModule,
		initializeHandlersModule,
	}

	return s
}

func (s *Server) Configure() error {
	log.Println("Configuring modules...")

	err := s.bootstrapper.Configure()

	if err != nil {
		return err
	}

	s.port = config.EnvVariables.App.Port

	log.Println("Setting up routes...")
	s.router = NewRouter()

	return nil
}

func (s *Server) ListenAndServe() error {
	return http.ListenAndServe(
		fmt.Sprintf(":%d", s.port),
		s.router,
	)
}
