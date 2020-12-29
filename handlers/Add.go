package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/woot-crdt/woot"
	log "github.com/sirupsen/logrus"
)

// addBody ...
type addBody struct {
	Value    string `json:"value"`
	Position int    `json:"position"`
}

// Add is the HTTP handler used to append
// values to the WString node in the server
func Add(w http.ResponseWriter, r *http.Request) {
	var err error
	var requestBody addBody
	var WStringPointer *woot.WString

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Add the given value to our stored WString
	WStringPointer, err = WString.GenerateInsert(requestBody.Position, requestBody.Value)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to add value")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Dereference the pointer to
	// the updated WString
	WString = *WStringPointer

	// DEBUG log in the case of success indicating
	// the new WString and the value added
	log.WithFields(log.Fields{
		"text":     woot.Value(WString),
		"value":    requestBody.Value,
		"position": requestBody.Position,
	}).Debug("successful wstring addition")

	// TODO: Broadcast Add

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
