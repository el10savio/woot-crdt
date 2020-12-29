package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/el10savio/woot-crdt/woot"
	log "github.com/sirupsen/logrus"
)

// deleteBody ...
type deleteBody struct {
	Position int `json:"position"`
}

// Delete is the HTTP handler used to delete
// values in the WString node in the server
func Delete(w http.ResponseWriter, r *http.Request) {
	var err error
	var requestBody deleteBody
	var WStringPointer *woot.WString

	// Obtain the value & position from POST Request Body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestBody)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed parse request body")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Delete the given value in our stored WString
	WStringPointer = WString.GenerateDelete(requestBody.Position)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to delete value")
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
		"position": requestBody.Position,
	}).Debug("successful wstring deletion")

	// Broadcast Delete
	err = BroadcastDelete(requestBody.Position)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("failed to sync delete value")
	}

	// Return HTTP 200 OK in the case of success
	w.WriteHeader(http.StatusOK)
}
