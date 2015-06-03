package routes

import (
	"errors"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func CalendarsGET(c *serve.Conn, db services.DB) {
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
	stringID := c.ParamVal("calendar_id")
	if stringID == "" {
		BadParam(c, "calendar_id")
		return
	}

	id, err := db.ParseID(stringID)
	if err != nil {
		BadParam(c, "calendar_id")
		return
	}
	// --- }}}

	// --- Find the Calendar {{{
	calendar := models.NewCalendar()
	calendar.SetID(id)

	if err := db.PopulateByID(calendar); err != nil {
		ServerError(c, err)
		return
	}

	if calendar.UserID != user.ID().String() {
		Unauthorized(c)
		return
	}
	// }}}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(calendar)),
	)
}

func CalendarsPOST(c *serve.Conn, db services.DB) {
}

func CalendarsDELETE(c *serve.Conn, db services.DB) {
}
