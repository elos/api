package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const FixtureIDParam = "fixture_id"

func retrieveFixture(c *serve.Conn, db services.DB) (*models.Fixture, bool) {
	id, ok := retrieveIDParam(FixtureIDParam, c, db)
	if !ok {
		return nil, false
	}

	fixture, err := models.FindFixture(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return fixture, true
}

// --- FixturesGET {{{
func FixturesGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	fixture, ok := retrieveFixture(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, fixture, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(fixture)),
	)
}

// --- }}}

// --- FixturesPOST {{{
func FixturesPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	fixture := models.NewFixture()
	if err := decoder.Decode(fixture); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if fixture.Id == "" {
		fixture.SetID(db.NewID())
		creation = true
	}

	if fixture.OwnerID == "" {
		fixture.OwnerID = user.Id
	}

	if user.ID().String() != fixture.OwnerID {
		Unauthorized(c)
		return
	}

	if err := db.Save(fixture); err != nil {
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
		transfer.StringMap(transfer.Map(fixture)),
	)
}

// --- }}}

func FixturesDELETE(c *serve.Conn, db services.DB) {
}

func FixturesOPTIONS(c *serve.Conn) {
	c.WriteHeader(200)
}
