package routes

import (
	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func CalendarsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	id, err := db.ParseID(c.ParamVal("calendar_id"))
	if err != nil {
		BadParam(c, "calendar_id")
		return
	}

	// --- Find the Calendar {{{
	calendar := models.NewCalendar()
	calendar.SetID(id)

	if err := db.PopulateByID(calendar); err != nil {
		ServerError(c, err)
		return
	}

	if calendar.OwnerID != user.ID().String() {
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
