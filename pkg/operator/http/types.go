package http

import (
	"regexp"
)

// AllocateArgs contains information necessary to allocate a GameServer
type AllocateArgs struct {
	SessionID      string   `json:"sessionID"`
	BuildID        string   `json:"buildID"`
	SessionCookie  string   `json:"sessionCookie"`
	InitialPlayers []string `json:"initialPlayers"`
}

// isValidUUID returns true if the string is a valid UUID
func isValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// validateAllocateArgs validates an instance of the AllocateArgs struct.
func validateAllocateArgs(aa *AllocateArgs) bool {
	if !isValidUUID(aa.SessionID) || !isValidUUID(aa.BuildID) {
		return false
	}
	return true
}

// RequestMultiplayerServerResponse contains details that are returned on a successful GameServer allocation call
type RequestMultiplayerServerResponse struct {
	IPV4Address string
	Ports       string
	SessionID   string
}
