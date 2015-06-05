package middleware

import (
	"errors"

	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const UserArtifact = "session-user"

type SessionAuth struct{}

func (sa *SessionAuth) Inbound(c *serve.Conn) bool {
	return true
}

func (sa *SessionAuth) Outbound(c *serve.Conn) bool {
	return true
}

// Retrieve's the UserArtifact produced by the SessionAuth
// middleware. Asserts a few common things, such as existence
// and type. Similiar pattern, u, ok := RetrieveUser(c).
// if !ok { return }
func RetrieveUser(c *serve.Conn, errorFn func(*serve.Conn, error)) (*models.User, bool) {
	v, ok := c.Context(UserArtifact)
	if !ok {
		errorFn(c, errors.New("User Artifact Missing"))
		return nil, false
	}
	user, ok := v.(*models.User)
	if !ok {
		errorFn(c, errors.New("User Cast Failed"))
		return nil, false
	}

	return user, true
}
