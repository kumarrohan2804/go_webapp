package main

import (
	"fmt"
	"net/http"
	"github.com/crud/handler"
)

const portNumber = ":8080"



func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Starting the application on port %s\n", portNumber)

	http.ListenAndServe(portNumber, nil)
}
