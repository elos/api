package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/elos/models/access"
)

// --- Helpers {{{
func retrieveModel(k data.Kind, c *serve.Conn, db services.DB) (data.Record, bool) {
	id, ok := retrieveIDParam(string(k)+"_id", c, db)
	if !ok {
		return nil, false
	}

	model := models.ModelFor(k)
	model.SetID(*id)

	if err := db.PopulateByID(model); err != nil {
		ServerError(c, err)
		return nil, false
	}

	return model, true
}

// --- }}}

// --- GET {{{
func GET(k data.Kind, c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	model, ok := retrieveModel(k, c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, model.(access.ModelProperty), c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(model)),
	)
}

// --- }}}

// --- POST {{{
func POST(k data.Kind, c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	decoder := json.NewDecoder(c.Request().Body)
	model := models.ModelFor(k)
	if err := decoder.Decode(model); err != nil {
		ServerError(c, err)
		return
	}

	creation := false

	if model.ID().String() == "" {
		model.SetID(db.NewID())
		creation = true
	}

	if !creation && !checkWriteAccess(user, model.(access.ModelProperty), c, db) {
		return
	}

	if err := db.Save(model); err != nil {
		ServerError(c, err)
		return
	}

	var status int
	if creation {
		status = 201
	} else {
		status = 200
	}

	c.Response(
		status,
		transfer.StringMap(transfer.Map(model)),
	)
}

// --- }}}

// --- DELETE {{{
func DELETE(k data.Kind, c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	model, ok := retrieveModel(k, c, db)
	if !ok {
		return
	}

	if !checkWriteAccess(user, model.(access.ModelProperty), c, db) {
		return
	}

	if err := db.Delete(model); err != nil {
		ServerError(c, err)
		return
	}

	c.Response(
		200,
		nil,
	)
}

// --- }}}

// --- OPTIONS {{{
func OPTIONS(k data.Kind, c *serve.Conn) {
	c.WriteHeader(200)
}

// --- }}}
