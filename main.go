package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	log.Printf("Serving HTTP on port %s...\n", port)
	log.Fatal((http.ListenAndServe(":"+port, nil)))
}
