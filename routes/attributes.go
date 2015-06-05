package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func AttributesGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	id, err := db.ParseID(c.ParamVal("attribute_id"))
	if err != nil {
		BadParam(c, "attribute_id")
		return
	}

	// --- Find the Attribute {{{
	attribute := models.NewAttribute()
	attribute.SetID(id)

	if err := db.PopulateByID(attribute); err != nil {
		ServerError(c, err)
		return
	}

	if attribute.UserID != user.ID().String() {
		Unauthorized(c)
		return
	}
	// }}}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(attribute)),
	)
}

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

	if user.ID().String() != attribute.UserID {
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

func AttributesDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Retrieve the ID {{{
	stringID := c.ParamVal("attribute_id")
	if stringID == "" {
		BadParam(c, "attribute_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "attribute_id")
		return
	}
	// --- }}}

	// --- Delete the Attribute {{{
	attribute := models.NewAttribute()
	attribute.SetID(id)

	if err := db.PopulateByID(attribute); err != nil {
		ServerError(c, err)
		return
	}

	if attribute.UserID != user.ID().String() {
		Unauthorized(c)
		return
	}

	if err := db.Delete(attribute); err != nil {
		ServerError(c, err)
		return
	}
	// }}}

	c.Response(
		200,
		nil,
	)
}
