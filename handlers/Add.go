package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Add is the HTTP handler used to append
// values to the WString node in the server
func Add(w http.ResponseWriter, r *http.Request) {
	var err error

	// Obtain the value & position from URL params
	value := mux.Vars(r)["value"]
	position := mux.Vars(r)["position"]

	// Add the given value to our stored WString
	WString, err = WString.GenerateInsert(position, value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// DEBUG log in the case of success indicating
	// the new WString and the value added
	log.WithFields(log.Fields{
		"text":     WString.Value(),
		"value":    value,
		"position": position,
	}).Debug("successful wstring addition")

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
