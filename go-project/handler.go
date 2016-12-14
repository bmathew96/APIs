package main

import (
	"fmt"
	"net/http"
)

func getHandler() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi there, I love %s!", r.)
	})

}
