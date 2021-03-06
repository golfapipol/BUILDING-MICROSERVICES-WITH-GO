package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 8000

	http.HandleFunc("/helloworld", helloWorldHandler)
	log.Printf("Server starting on port %v", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}
