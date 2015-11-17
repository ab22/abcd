package main

import (
	"log"

	"github.com/ab22/abcd/config"
	"github.com/ab22/abcd/handlers"
	"github.com/ab22/abcd/services"
)

// bootstrapFunc is a simple function that configures the app modules, takes
// no params and returns an error if anything goes wrong.
type bootstrapFunc func() error

// boostrapper contains all of the bootstrapper functions to be executed.
type bootstrapper []bootstrapFunc

// Configure executes the bootstrapper functions. If there's an error, it
// stops iterating over the functions and returns the error.
func (b bootstrapper) Configure() error {
	var err error

	for _, f := range b {
		if err = f(); err != nil {
			return err
		}
	}

	return nil
}

// Configure abcd/config package.
func initializeConfigurationModule() error {
	log.Println("Loading config...")

	if err := config.Initialize(); err != nil {
		return err
	}

	config.EnvVariables.Print()
	return nil
}

// Configure playcheck/services package.
func initializeServicesModule() error {
	log.Println("Initializing services...")

	if err := services.Initialize(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

// Configure abcd/handlers package.
func initializeHandlersModule() error {
	log.Println("Initializing handlers...")
	handlers.Initialize()

	return nil
}
