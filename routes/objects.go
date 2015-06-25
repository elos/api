package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func ObjectsGET(c *serve.Conn, db services.DB) {
	GET(models.ObjectKind, c, db)
}

func ObjectsPOST(c *serve.Conn, db services.DB) {
	POST(models.ObjectKind, c, db)
}

func ObjectsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ObjectKind, c, db)
}

func ObjectsOPTIONS(c *serve.Conn) {
	OPTIONS(models.ObjectKind, c)
}
