package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func ClassesGET(c *serve.Conn, db services.DB) {
	GET(models.ClassKind, c, db)
}

func ClassesPOST(c *serve.Conn, db services.DB) {
	POST(models.ClassKind, c, db)
}

func ClassesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ClassKind, c, db)
}

func ClassesOPTIONS(c *serve.Conn) {
	OPTIONS(models.ClassKind, c)
}
