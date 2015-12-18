package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ab22/abcd/config"
)

var (
	// Defines all of the functions to be executed and configured on boot.
	abcdBootstrapper = bootstrapper{
		initializeConfigurationModule,
		initializeServicesModule,
		initializeHandlersModule,
	}
)

func main() {
	var (
		err    error
		port   int
		router *Router
	)

	log.Println("Starting server...")

	if err = abcdBootstrapper.Configure(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Setting up routes...")

	router = NewRouter()

	log.Println("Listening...")

	port = config.EnvVariables.App.Port
	err = http.ListenAndServe(
		fmt.Sprintf(":%d", port),
		router,
	)

	if err != nil {
		log.Fatalln(err)
	}
}
