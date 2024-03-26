package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create a file server handler to serve files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define a handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("index.html")

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, string(data))
	}

	// Register the handler function with the default serve mux (router)
	http.HandleFunc("/", handler)

	log.Println("Listening on port 5000")
	http.ListenAndServe(":5000", nil)
}
