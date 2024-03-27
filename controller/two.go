package controller

import (
	"net/http"

	"github.com/google/martian/log"
)

func TwoHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("Two Handler")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Two Handler"))
}
