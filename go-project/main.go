package main

import (
	"log"
	"net/http"
	"time"
)

var config EnvConfig

func init() {

	log.Println("Grabbing Environment Variables")
	config, err := getEnvConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	getHandler()
	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Serving on port 8080")
	log.Fatal(s.ListenAndServe())

}
