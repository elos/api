package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
)

func UsersGET(c *serve.Conn, db services.DB) {}

func UsersPOST(c *serve.Conn, db services.DB) {}

func UsersDELETE(c *serve.Conn, db services.DB) {}
