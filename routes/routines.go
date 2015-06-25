package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func RoutinesGET(c *serve.Conn, db services.DB) {
	GET(models.RoutineKind, c, db)
}

func RoutinesPOST(c *serve.Conn, db services.DB) {
	POST(models.RoutineKind, c, db)
}

func RoutinesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.RoutineKind, c, db)
}

func RoutinesOPTIONS(c *serve.Conn) {
	OPTIONS(models.RoutineKind, c)
}
