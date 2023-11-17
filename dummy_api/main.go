package main

import (
	swagger "dummy_api/go"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define a command-line flag for the port
	port := flag.Int("port", 8080, "Port for the server")
	flag.Parse()

	// Print the selected port
	fmt.Printf("Server started on port %d\n", *port)

	router := swagger.NewRouter()

	// Use the selected port for the server
	addr := fmt.Sprintf(":%d", *port)
	log.Fatal(http.ListenAndServe(addr, router))
}
