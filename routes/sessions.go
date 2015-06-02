package routes

import (
	"errors"

	"github.com/elos/app/middleware"
	"github.com/elos/app/services"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/elos/models/user"
)

func SessionsGET(c *serve.Conn, db services.DB) {
	// --- Retrieve the User {{{
	v, ok := c.Context(middleware.UserArtifact)
	if !ok {
		ServerError(c, errors.New("User Artifact Missing"))
		return
	}
	user, ok := v.(*models.User)
	if !ok {
		ServerError(c, errors.New("User Cast Failed"))
		return
	}
	// --- }}}

	// --- Retrieve the ID {{{
	stringID := c.ParamVal("session_id")
	if stringID == "" {
		BadParam(c, "session_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "session_id")
		return
	}
	// --- }}}

	// --- Find the Session {{{
	session := models.NewSession()
	session.SetID(id)

	if err := db.PopulateByID(session); err != nil {
		ServerError(c, err)
		return
	}

	if id.String() != user.ID().String() {
		Unauthorized(c)
		return
	}
	// }}}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(session)),
	)
}

func SessionsPOST(c *serve.Conn, db services.DB) {
	// --- Retrieve Credentials {{{
	credentials, err := c.ParamVals("id", "key")
	if err != nil {
		switch err.(type) {
		case *ehttp.MissingParamError:
			if string(*err.(*ehttp.MissingParamError)) == "id" {
				BadParam(c, "id")
				return
			} else {
				BadParam(c, "key")
				return
			}

			ServerError(c, err)
			return
		}
	}
	// --- }}}

	// --- Authenticate the user {{{
	u, err := user.Authenticate(db, credentials["id"], credentials["key"])
	if err != nil {
		if err.Error() == "invalid key" {
			Unauthorized(c)
			return
		}

		if err == data.ErrNotFound {
			RecordNotFound(c)
			return
		}

		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Create the session {{{
	session := models.NewSessionForUser(u)
	if err := db.Save(session); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(session)),
	)
}

func SessionsDELETE(c *serve.Conn, db services.DB) {
	// --- Retrieve the User {{{
	v, ok := c.Context(middleware.UserArtifact)
	if !ok {
		ServerError(c, errors.New("User Artifact Missing"))
		return
	}
	user, ok := v.(*models.User)
	if !ok {
		ServerError(c, errors.New("User Cast Failed"))
		return
	}
	// --- }}}

	// --- Retrieve the ID {{{
	stringID := c.ParamVal("session_id")
	if stringID == "" {
		BadParam(c, "session_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "session_id")
		return
	}
	// --- }}}

	// --- Find the Session {{{
	session := models.NewSession()
	session.SetID(id)

	if err := db.PopulateByID(session); err != nil {
		ServerError(c, err)
		return
	}

	if id.String() != user.ID().String() {
		Unauthorized(c)
		return
	}
	// }}}

	// --- Delete it {{{
	if err := db.Delete(session); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	c.Response(
		200,
		nil,
	)
}
