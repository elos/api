package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const ActionIDParam = "action_id"

func retrieveAction(c *serve.Conn, db services.DB) (*models.Action, bool) {
	id, ok := retrieveIDParam(ActionIDParam, c, db)
	if !ok {
		return nil, false
	}

	action, err := models.FindAction(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return action, true
}

// --- ActionsGET {{{

func ActionsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	action, ok := retrieveAction(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, action, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(action)),
	)
}

// --- }}}

// --- ActionsPOST {{{

func ActionsPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	action := models.NewAction()
	if err := decoder.Decode(action); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if action.Id == "" {
		action.SetID(db.NewID())
		creation = true
	}

	if user.ID().String() != action.OwnerID {
		Unauthorized(c)
		return
	}

	if err := db.Save(action); err != nil {
		ServerError(c, err)
		return
	}

	var status int
	if creation {
		status = 201
	} else {
		status = 200
	}
	// --- }}}

	c.Response(
		status,
		transfer.StringMap(transfer.Map(action)),
	)
}

// --- }}}

// --- ActionsDELETE {{{

func ActionsDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	action, ok := retrieveAction(c, db)
	if !ok {
		return
	}

	if !checkWriteAccess(user, action, c, db) {
		return
	}

	if err := db.Delete(action); err != nil {
		ServerError(c, err)
		return
	}

	c.Response(
		200,
		nil,
	)
}

// --- }}}
