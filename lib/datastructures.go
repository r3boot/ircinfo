package lib

/*
 * This file contains various data structures which are used throughout this
 * application.
 */

// Entity
type Entity struct {
	Mask     string
	NickName string
	Username string
	Hostname string
}

// List of entities bound to a certain host
type HostToEntity struct {
	Hostname string
	Entity   []string
}

// List of entities bound to a certain nickname
type NickToEntity struct {
	NickName string
	Entity   []string
}

// List of entities bound to a certain username
type UserToEntity struct {
	Username string
	Entity   []string
}

// Record nickchanges
type NickChange struct {
	OldMask string
	NewMask string
}

// Constants used throughout the application
const CMD_SHUTDOWN int = 0 // When sent to a goroutine, shut it down
