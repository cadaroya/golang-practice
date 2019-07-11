package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Routes: URL mapped to function
	http.HandleFunc("/hello", helloRequest)
	http.HandleFunc("/", getRequest)

	// Start server in port 8000
	log.Println("Listening on http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// basic handler for /hello request
func helloRequest(w http.ResponseWriter, r *http.Request) {

	// Fprint writes to a writer, in this case the http response
	fmt.Fprint(w, "Hello. This is the server response for helloRequest :)")
	return
}

// this function simply serves up the server.go file
func getRequest(w http.ResponseWriter, r *http.Request) {
	fileRequested := "./" + r.URL.Path
	http.ServeFile(w, r, fileRequested)
	return
}
