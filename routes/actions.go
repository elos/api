package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func ActionsGET(c *serve.Conn, db services.DB) {
	GET(models.ActionKind, c, db)
}

func ActionsPOST(c *serve.Conn, db services.DB) {
	POST(models.ActionKind, c, db)
}

func ActionsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ActionKind, c, db)
}

func ActionsOPTIONS(c *serve.Conn) {
	OPTIONS(models.ActionKind, c)
}
