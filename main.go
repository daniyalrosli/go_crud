package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
