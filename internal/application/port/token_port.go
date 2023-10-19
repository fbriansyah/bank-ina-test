package port

import (
	"time"

	dmtoken "github.com/fbriansyah/bank-ina-test/internal/application/domain/token"
)

// Maker is an interface for managing tokens
type TokenMakerPort interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(mode string, userID int32, duration time.Duration) (string, *dmtoken.Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*dmtoken.Payload, error)
}
