package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func CalendarsGET(c *serve.Conn, db services.DB) {
	GET(models.CalendarKind, c, db)
}

func CalendarsPOST(c *serve.Conn, db services.DB) {
	POST(models.CalendarKind, c, db)
}

func CalendarsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.CalendarKind, c, db)
}

func CalendarsOPTIONS(c *serve.Conn) {
	OPTIONS(models.CalendarKind, c)
}
