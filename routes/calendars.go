package routes

import (
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

func CalendarsPOST(c *serve.Conn, db services.DB) {
}

func CalendarsDELETE(c *serve.Conn, db services.DB) {
}
