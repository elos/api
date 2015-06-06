package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const AttributeIDParam = "attribute_id"

func retrieveAttribute(c *serve.Conn, db services.DB) (*models.Attribute, bool) {
	id, ok := retrieveIDParam(AttributeIDParam, c, db)
	if !ok {
		return nil, false
	}

	attribute, err := models.FindAttribute(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return attribute, true
}

// --- AttributesGET {{{

func AttributesGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	attribute, ok := retrieveAttribute(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, attribute, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(attribute)),
	)
}

// --- }}}

// --- AttributesPOST {{{

func AttributesPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	attribute := models.NewAttribute()
	if err := decoder.Decode(attribute); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if attribute.Id == "" {
		attribute.SetID(db.NewID())
		creation = true
	}

	if err := db.Save(attribute); err != nil {
		ServerError(c, err)
		return
	}

	if user.ID().String() != attribute.OwnerID {
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
		transfer.StringMap(transfer.Map(attribute)),
	)
}

// --- }}}

// --- AttributesDELETE {{{

func AttributesDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	attribute, ok := retrieveAction(c, db)
	if !ok {
		return
	}

	if !checkWriteAccess(user, attribute, c, db) {
		return
	}

	c.Response(
		200,
		nil,
	)
}

// --- }}}
