package token

import "time"

// Maker is an interface for manging tokens
type Maker interface {
	// CreateToken creates a new token a specific username and duration
	CreateToken(username string, duration time.Duration) (string, *Payload, error)

	// VerifyToken Check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
