package middleware

import (
	"errors"

	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const (
	// The HTTP Header used to specifiy authentication with the Elos API
	AuthHeader = "Elos-Auth"

	// The session associated with the token found in the Elos-Auth header.
	// Artifacts are a concept from the aeolus package, read more about them there
	SessionArtifact = "session"

	// The user associated with the session found in the SessionArtifact
	UserArtifact = "session-user"
)

// The structure encapsulating the necessary context for the operation
// of session authentication. Need a database access and ability to reject
// request. Satisfies the serve.Middleware interface.
type SessionAuth struct {
	// The database service required by the API, used for looking up
	// the session and the user
	DB services.DB

	// The handler used by this middleware to reject a request
	UnauthorizedHandler serve.Route
}

func (sa *SessionAuth) Inbound(c *serve.Conn) bool {
	// Retrieve the request's header
	header := c.Request().Header

	// Retrieve the Elos-Auth header
	auth, ok := header[AuthHeader]
	if !ok { // No AuthHeader
		sa.UnauthorizedHandler(c)
		return false
	}

	// You technically provide an array of values for all HTTP Headers,
	// so we are saying that we only accept when you provide one string
	// in the Elos-Auth header
	if len(auth) != 1 { // very strict, because they could have: ['validToken', 'foobar']
		sa.UnauthorizedHandler(c)
		return false
	}

	// Now get that first token
	token := auth[0]

	// Retrieve the session
	session, err := models.SessionForToken(sa.DB, token)
	if err != nil {
		// we could have failed (ServerError) but we have no obligation to tell anyone whose
		// identity we are unsure of about that
		sa.UnauthorizedHandler(c)
		return false
	}

	// Retrieve the session's owner (a models.User)
	// and a "user" in the Elos Access System
	user, err := session.Owner(sa.DB)
	if err != nil {
		sa.UnauthorizedHandler(c)
		return false
	}

	// Add context to the serve.Conn
	c.AddContext(SessionArtifact, session)
	c.AddContext(UserArtifact, user)

	return true // continues the request down the pipeline
}

func (sa *SessionAuth) Outbound(c *serve.Conn) bool {
	// currently, nothing to do after the request is written to (outbound middleware)
	return true
}

// Retrieve's the UserArtifact produced by the SessionAuth
// middleware. Asserts a few common things, such as existence
// and type. The pattern for using this function is
//		u, ok := RetrieveUser(c, ServerErrorHandler)
//      if !ok {
//			// no need to write to response, because if this failed something is fundamentally
//			// logically inconsistent in your application.
//			return
//      }
// You should only use this in a handler in which is wrapped by the SessionAuth middleware.
func RetrieveUser(c *serve.Conn, errorFn func(*serve.Conn, error)) (*models.User, bool) {
	// Retreive the UserArtiface value
	v, ok := c.Context(UserArtifact)
	if !ok {
		errorFn(c, errors.New("User Artifact Missing"))
		return nil, false
	}

	// Assert type
	user, ok := v.(*models.User)
	if !ok {
		errorFn(c, errors.New("User Cast Failed"))
		return nil, false
	}

	return user, true
}
