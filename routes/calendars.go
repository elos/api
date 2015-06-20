package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const CalendarIDParam = "calendar_id"

func retrieveCalendar(c *serve.Conn, db services.DB) (*models.Calendar, bool) {
	id, ok := retrieveIDParam(CalendarIDParam, c, db)
	if !ok {
		return nil, false
	}

	calendar, err := models.FindCalendar(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return calendar, true
}

// --- CalendarsGET {{{

func CalendarsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	calendar, ok := retrieveCalendar(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, calendar, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(calendar)),
	)
}

// --- }}}

// --- CalendarsPOST {{{
func CalendarsPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	calendar := models.NewCalendar()
	if err := decoder.Decode(calendar); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if calendar.Id == "" {
		calendar.SetID(db.NewID())
		creation = true
	}

	if user.ID().String() != calendar.OwnerID {
		Unauthorized(c)
		return
	}

	if err := db.Save(calendar); err != nil {
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
		transfer.StringMap(transfer.Map(calendar)),
	)
}

// --- }}}

// --- CalendarsDELETE {{{
func CalendarsDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	calendar, ok := retrieveCalendar(c, db)
	if !ok {
		return
	}

	if !checkWriteAccess(user, calendar, c, db) {
		return
	}

	if err := db.Delete(calendar); err != nil {
		ServerError(c, err)
		return
	}

	c.Response(
		200,
		nil,
	)
}

// --- }}}
