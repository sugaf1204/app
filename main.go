package main

import (
	"fmt"
	"net/http"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})
	fmt.Println("Starting server on port :8080")

	http.ListenAndServe(":8080", r)
}
