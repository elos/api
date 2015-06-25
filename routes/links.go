package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func LinksGET(c *serve.Conn, db services.DB) {
	GET(models.LinkKind, c, db)
}

func LinksPOST(c *serve.Conn, db services.DB) {
	POST(models.LinkKind, c, db)
}

func LinksDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.LinkKind, c, db)
}

func LinksOPTIONS(c *serve.Conn) {
	OPTIONS(models.LinkKind, c)
}
