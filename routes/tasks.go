package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func TasksGET(c *serve.Conn, db services.DB) {
	GET(models.TaskKind, c, db)
}

func TasksPOST(c *serve.Conn, db services.DB) {
	POST(models.TaskKind, c, db)
}

func TasksDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.TaskKind, c, db)
}

func TasksOPTIONS(c *serve.Conn, db services.DB) {
	OPTIONS(models.TaskKind, c)
}
