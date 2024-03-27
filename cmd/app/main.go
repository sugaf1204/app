package main

import (
	"net/http"

	"github.com/sugaf1204/app/controller"

	"github.com/google/martian/log"
)

func main() {

	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("Default Handler")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})

	r.HandleFunc("/one", controller.OneHandler)
	r.HandleFunc("/two", controller.TwoHandler)

	log.Infof("Starting server on port :8080")

	http.ListenAndServe(":8080", r)
}
