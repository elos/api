package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func EventsGET(c *serve.Conn, db services.DB) {
	GET(models.EventKind, c, db)
}

func EventsPOST(c *serve.Conn, db services.DB) {
	POST(models.EventKind, c, db)
}

func EventsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.EventKind, c, db)
}

func EventsOPTIONS(c *serve.Conn) {
	OPTIONS(models.EventKind, c)
}
