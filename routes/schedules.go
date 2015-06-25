package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func SchedulesGET(c *serve.Conn, db services.DB) {
	GET(models.ScheduleKind, c, db)
}

func SchedulesPOST(c *serve.Conn, db services.DB) {
	POST(models.ScheduleKind, c, db)
}

func SchedulesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ScheduleKind, c, db)
}

func SchedulesOPTIONS(c *serve.Conn) {
	OPTIONS(models.ScheduleKind, c)
}
