package middleware

import (
	"errors"

	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const (
	AuthHeader   = "Elos-Auth"
	UserArtifact = "session-user"
)

type SessionAuth struct {
	DB                  services.DB
	UnauthorizedHandler serve.Route
}

func (sa *SessionAuth) Inbound(c *serve.Conn) bool {
	header := c.Request().Header
	auth, ok := header[AuthHeader]
	if !ok {
		sa.UnauthorizedHandler(c)
		return false
	}

	// be very strict about information provided
	if len(auth) != 1 {
		sa.UnauthorizedHandler(c)
		return false
	}

	token := auth[0]

	session, err := models.SessionForToken(sa.DB, token)
	if err != nil {
		sa.UnauthorizedHandler(c)
		return false
	}

	user, err := session.Owner(sa.DB)
	if err != nil {
		sa.UnauthorizedHandler(c)
		return false
	}

	c.AddContext(UserArtifact, user)

	return true
}

func (sa *SessionAuth) Outbound(c *serve.Conn) bool {
	// nothing to do
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
