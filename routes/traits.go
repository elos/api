package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
)

func TraitsGET(c *serve.Conn, db services.DB)    {}
func TraitsPOST(c *serve.Conn, db services.DB)   {}
func TraitsDELETE(c *serve.Conn, db services.DB) {}
