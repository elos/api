package routes

import (
	"encoding/json"
	"errors"

	"github.com/elos/api/services"
	"github.com/elos/app/middleware"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func ActionsGET(c *serve.Conn, db services.DB) {
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
	stringID := c.ParamVal("action_id")
	if stringID == "" {
		BadParam(c, "action_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "action_id")
		return
	}
	// --- }}}

	// --- Find the Action {{{
	action := models.NewAction()
	action.SetID(id)

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
	stringID := c.ParamVal("action_id")
	if stringID == "" {
		BadParam(c, "action_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "action_id")
		return
	}
	// --- }}}

	// --- Delete the Action {{{
	action := models.NewAction()
	action.SetID(id)

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
