package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// BroadcastAdd ...
func BroadcastAdd(position int, value string) error {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return an error if no peers are present
	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	// Iterate over the peer list and send a
	// /woot/sync/add POST request to each peer
	for _, peer := range peers {
		err := SendAddRequest(peer, addBody{Value: value, Position: position})
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending woot add broadcast")
			continue
		}
	}

	// DEBUG log in the case of success
	log.WithFields(log.Fields{
		"value":    value,
		"position": position,
	}).Debug("successful woot add broadcast")

	return nil
}

// SendAddRequest handles sending of /woot/sync/add HTTP POST Request
func SendAddRequest(peer string, postBody addBody) error {
	// Return an error if the peer is nil
	if peer == "" {
		return errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/woot/sync/add", peer, GetNetwork())

	json, _ := json.Marshal(postBody)
	encodedBody := bytes.NewReader(json)

	request, err := http.NewRequest("POST", url, encodedBody)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Duration(5 * 60 * time.Second),
	}

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// BroadcastDelete ...
func BroadcastDelete(position int) error {
	// Obtain addresses of peer nodes in the cluster
	peers := GetPeerList()

	// Return an error if no peers are present
	if len(peers) == 0 {
		return errors.New("nil peers present")
	}

	// Iterate over the peer list and send a
	// /woot/sync/delete POST request to each peer
	for _, peer := range peers {
		err := SendDeleteRequest(peer, deleteBody{Position: position})
		if err != nil {
			log.WithFields(log.Fields{"error": err, "peer": peer}).Error("failed sending woot delete broadcast")
			continue
		}
	}

	// DEBUG log in the case of success
	log.WithFields(log.Fields{
		"position": position,
	}).Debug("successful woot delete broadcast")

	return nil
}

// SendDeleteRequest handles sending of /woot/sync/delete HTTP POST Request
func SendDeleteRequest(peer string, postBody deleteBody) error {
	// Return an error if the peer is nil
	if peer == "" {
		return errors.New("empty peer provided")
	}

	// Resolve the Peer ID and network to generate the request URL
	url := fmt.Sprintf("http://%s.%s/woot/sync/delete", peer, GetNetwork())

	json, _ := json.Marshal(postBody)
	encodedBody := bytes.NewReader(json)

	request, err := http.NewRequest("POST", url, encodedBody)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Duration(5 * 60 * time.Second),
	}

	response, err := client.Do(request)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}
