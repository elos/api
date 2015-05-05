package api

import (
	"github.com/elos/api/hermes"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

// THIS FILE GENERATED USING METIS
// TO REGENERATE BASED ON NEW DEFINITIONS
// cd gen/ && go run gen.go

func router(sockets autonomous.Manager, db data.DB) serve.Router {
	r := builtin.NewRouter()

	r.GET("/action", Serve(hermes.GET, "action", db))
	r.POST("/action", Serve(hermes.POST, "action", db))
	r.DELETE("/action", Serve(hermes.DELETE, "action", db))

	r.GET("/attribute", Serve(hermes.GET, "attribute", db))
	r.POST("/attribute", Serve(hermes.POST, "attribute", db))
	r.DELETE("/attribute", Serve(hermes.DELETE, "attribute", db))

	r.GET("/calendar", Serve(hermes.GET, "calendar", db))
	r.POST("/calendar", Serve(hermes.POST, "calendar", db))
	r.DELETE("/calendar", Serve(hermes.DELETE, "calendar", db))

	r.GET("/class", Serve(hermes.GET, "class", db))
	r.POST("/class", Serve(hermes.POST, "class", db))
	r.DELETE("/class", Serve(hermes.DELETE, "class", db))

	r.GET("/event", Serve(hermes.GET, "event", db))
	r.POST("/event", Serve(hermes.POST, "event", db))
	r.DELETE("/event", Serve(hermes.DELETE, "event", db))

	r.GET("/fixture", Serve(hermes.GET, "fixture", db))
	r.POST("/fixture", Serve(hermes.POST, "fixture", db))
	r.DELETE("/fixture", Serve(hermes.DELETE, "fixture", db))

	r.GET("/link", Serve(hermes.GET, "link", db))
	r.POST("/link", Serve(hermes.POST, "link", db))
	r.DELETE("/link", Serve(hermes.DELETE, "link", db))

	r.GET("/object", Serve(hermes.GET, "object", db))
	r.POST("/object", Serve(hermes.POST, "object", db))
	r.DELETE("/object", Serve(hermes.DELETE, "object", db))

	r.GET("/ontology", Serve(hermes.GET, "ontology", db))
	r.POST("/ontology", Serve(hermes.POST, "ontology", db))
	r.DELETE("/ontology", Serve(hermes.DELETE, "ontology", db))

	r.GET("/relation", Serve(hermes.GET, "relation", db))
	r.POST("/relation", Serve(hermes.POST, "relation", db))
	r.DELETE("/relation", Serve(hermes.DELETE, "relation", db))

	r.GET("/routine", Serve(hermes.GET, "routine", db))
	r.POST("/routine", Serve(hermes.POST, "routine", db))
	r.DELETE("/routine", Serve(hermes.DELETE, "routine", db))

	r.GET("/schedule", Serve(hermes.GET, "schedule", db))
	r.POST("/schedule", Serve(hermes.POST, "schedule", db))
	r.DELETE("/schedule", Serve(hermes.DELETE, "schedule", db))

	r.GET("/task", Serve(hermes.GET, "task", db))
	r.POST("/task", Serve(hermes.POST, "task", db))
	r.DELETE("/task", Serve(hermes.DELETE, "task", db))

	r.GET("/trait", Serve(hermes.GET, "trait", db))
	r.POST("/trait", Serve(hermes.POST, "trait", db))
	r.DELETE("/trait", Serve(hermes.DELETE, "trait", db))

	r.GET("/user", Serve(hermes.GET, "user", db))
	r.POST("/user", Serve(hermes.POST, "user", db))
	r.DELETE("/user", Serve(hermes.DELETE, "user", db))

	return r
}
