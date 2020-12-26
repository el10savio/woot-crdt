package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/woot-crdt/woot"
	log "github.com/sirupsen/logrus"
)

// List is the HTTP handler used to return
// all the values present in the WString node in the server
func List(w http.ResponseWriter, r *http.Request) {
	// Get the values from the WString
	text := woot.Value(WString)

	// DEBUG log in the case of success
	// indicating the new WString
	log.WithFields(log.Fields{
		"text": text,
	}).Debug("successful wstring list")

	// JSON encode response value
	json.NewEncoder(w).Encode(text)
}
