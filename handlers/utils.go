package handlers

import (
	"os"
	"strings"
)

// GetPeerList Obtains Peer List
// From Environment Variable
func GetPeerList() []string {
	if os.Getenv("PEERS") == "" {
		return []string{}
	}
	return strings.Split(os.Getenv("PEERS"), ",")
}

// GetNetwork Obtains Network
// From Environment Variable
func GetNetwork() string {
	return os.Getenv("NETWORK") + ":8080"
}
