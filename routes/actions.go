package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func retrieveActionID(c *serve.Conn, db services.DB) (*data.ID, bool) {
	id, err := db.ParseID(c.ParamVal("action_id"))
	if err != nil {
		BadParam(c, "action_id")
		return new(data.ID), false
	}

	return &id, true
}

func ActionsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	id, ok := retrieveActionID(c, db)
	if !ok {
		return
	}

	// --- Find the Action {{{
	action := models.NewAction()
	action.SetID(*id)

	if err := db.PopulateByID(action); err != nil {
		ServerError(c, err)
		return
	}

	if action.UserID != user.ID().String() {
		Unauthorized(c)
		return
	}
	// }}}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(action)),
	)
}

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

	if err := db.Save(action); err != nil {
		ServerError(c, err)
		return
	}

	if user.ID().String() != action.UserID {
		Unauthorized(c)
		return
	}

	var status uint64
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

func ActionsDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	id, ok := retrieveActionID(c, db)
	if !ok {
		return
	}

	// --- Delete the Action {{{
	action := models.NewAction()
	action.SetID(*id)

	if err := db.PopulateByID(action); err != nil {
		ServerError(c, err)
		return
	}

	if action.UserID != user.ID().String() {
		Unauthorized(c)
		return
	}

	if err := db.Delete(action); err != nil {
		ServerError(c, err)
		return
	}
	// }}}

	c.Response(
		200,
		nil,
	)
}
