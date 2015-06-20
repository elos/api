package routes

import (
	"errors"
	"time"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const SessionIDParam = "session_id"

func retrieveSession(c *serve.Conn, db services.DB) (*models.Session, bool) {
	id, ok := retrieveIDParam(SessionIDParam, c, db)
	if !ok {
		return nil, false
	}

	session, err := models.FindSession(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return session, true
}

// --- SessionsGET {{{

func SessionsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Specialty, if no session specified, return current {{{
	if c.ParamVal(SessionIDParam) == "" {
		v, ok := c.Context(middleware.SessionArtifact)
		if !ok {
			ServerError(c, errors.New("Session Artifact Missing"))
			return
		}

		session, ok := v.(*models.Session)
		if !ok {
			ServerError(c, errors.New("Session Cast Failed"))
			return
		}

		c.Response(
			200,
			transfer.StringMap(transfer.Map(session)),
		)

		return
	}
	// --- }}}

	session, ok := retrieveSession(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, session, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(session)),
	)
}

// --- }}}

// --- SessionsPOST {{{
func SessionsPOST(c *serve.Conn, db services.DB) {
	credentials, err := c.ParamVals("public", "private")
	if err != nil {
		switch err.(type) {
		case *ehttp.MissingParamError:
			if string(*err.(*ehttp.MissingParamError)) == "public" {
				BadParam(c, "public")
				return
			} else {
				BadParam(c, "private")
				return
			}

			ServerError(c, err)
			return
		}
	}

	credential, err := models.Authenticate(db, credentials["public"], credentials["private"])

	if err != nil {
		Unauthorized(c)
		return
	}

	session, err := credential.NewSession(db, 3600*time.Second)
	if err != nil {
		ServerError(c, err)
	}

	c.Response(
		201,
		transfer.StringMap(transfer.Map(session)),
	)
}

// --- }}}

// --- SessionsDELETE {{{

func SessionsDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	session, ok := retrieveSession(c, db)
	if !ok {
		return
	}

	if !checkWriteAccess(user, session, c, db) {
		return
	}

	if err := db.Delete(session); err != nil {
		ServerError(c, err)
		return
	}

	c.Response(
		200,
		nil,
	)
}

// --- }}}

func SessionsOPTIONS(c *serve.Conn) {
	c.Write([]byte(""))
}
