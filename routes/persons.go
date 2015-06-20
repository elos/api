package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const PersonIDParam = "person_id"

func retrievePerson(c *serve.Conn, db services.DB) (*models.Person, bool) {
	id, ok := retrieveIDParam(PersonIDParam, c, db)
	if !ok {
		return nil, false
	}

	person, err := models.FindPerson(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return person, true
}

// --- PersonsGET {{{
func PersonsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	if c.ParamVal(PersonIDParam) == "" {
		person := models.NewPerson()

		if err := db.PopulateByField("owner_id", user.ID().String(), person); err != nil {
			ServerError(c, err)
			return
		}

		c.Response(
			200,
			transfer.StringMap(transfer.Map(person)),
		)

		return
	}

	person, ok := retrievePerson(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, person, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(person)),
	)
}

// --- }}}

// --- PersonsPOST {{{
func PersonsPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	person := models.NewPerson()
	if err := decoder.Decode(person); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if person.Id == "" {
		person.SetID(db.NewID())
		creation = true
	}

	if user.ID().String() != person.OwnerID {
		Unauthorized(c)
		return
	}

	if err := db.Save(person); err != nil {
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
		transfer.StringMap(transfer.Map(person)),
	)
}

// --- }}}

func PersonsDELETE(c *serve.Conn, db services.DB) {
}

func PersonsOPTIONS(c *serve.Conn) {
	c.WriteHeader(200)
}
