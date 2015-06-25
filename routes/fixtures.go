package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func FixturesGET(c *serve.Conn, db services.DB) {
	GET(models.FixtureKind, c, db)
}

func FixturesPOST(c *serve.Conn, db services.DB) {
	POST(models.FixtureKind, c, db)
}

func FixturesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.FixtureKind, c, db)
}

func FixturesOPTIONS(c *serve.Conn) {
	OPTIONS(models.FixtureKind, c)
}
