package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func RelationsGET(c *serve.Conn, db services.DB) {
	GET(models.RelationKind, c, db)
}

func RelationsPOST(c *serve.Conn, db services.DB) {
	POST(models.RelationKind, c, db)
}

func RelationsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.RelationKind, c, db)
}

func RelationsOPTIONS(c *serve.Conn) {
	OPTIONS(models.RelationKind, c)
}
