package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func OntologiesGET(c *serve.Conn, db services.DB) {
	GET(models.OntologyKind, c, db)
}

func OntologiesPOST(c *serve.Conn, db services.DB) {
	POST(models.OntologyKind, c, db)
}

func OntologiesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.OntologyKind, c, db)
}

func OntologiesOPTIONS(c *serve.Conn) {
	OPTIONS(models.OntologyKind, c)
}
