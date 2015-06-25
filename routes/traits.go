package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func TraitsGET(c *serve.Conn, db services.DB) {
	GET(models.TraitKind, c, db)
}
func TraitsPOST(c *serve.Conn, db services.DB) {
	POST(models.TraitKind, c, db)
}
func TraitsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.TraitKind, c, db)
}
func TraitsOPTIONS(c *serve.Conn) {
	OPTIONS(models.TraitKind, c)
}
