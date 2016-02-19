package main

import (
	"log"
)

func main() {
	var (
		err error
		s   *Server
	)

	log.Println("Starting server...")

	s = NewServer()

	log.Println("Listening...")
	err = s.ListenAndServe()

	if err != nil {
		log.Fatalln(err)
	}
}
