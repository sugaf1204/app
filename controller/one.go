package controller

import (
	"net/http"

	"github.com/google/martian/log"
)

func OneHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("One Handler")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("One Handler"))
}
