package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func AttributesGET(c *serve.Conn, db services.DB) {
	GET(models.AttributeKind, c, db)
}

func AttributesPOST(c *serve.Conn, db services.DB) {
	POST(models.AttributeKind, c, db)
}

func AttributesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.AttributeKind, c, db)
}

func AttributesOPTIONS(c *serve.Conn) {
	OPTIONS(models.AttributeKind, c)
}
