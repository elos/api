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

func router(db data.DB, sockets autonomous.Manager) serve.Router {
	r := builtin.NewRouter()

	r.GET("/actions", Serve(hermes.GET, "action", db))
	r.POST("/actions", Serve(hermes.POST, "action", db))
	r.DELETE("/actions", Serve(hermes.DELETE, "action", db))

	r.GET("/attributes", Serve(hermes.GET, "attribute", db))
	r.POST("/attributes", Serve(hermes.POST, "attribute", db))
	r.DELETE("/attributes", Serve(hermes.DELETE, "attribute", db))

	r.GET("/calendars", Serve(hermes.GET, "calendar", db))
	r.POST("/calendars", Serve(hermes.POST, "calendar", db))
	r.DELETE("/calendars", Serve(hermes.DELETE, "calendar", db))

	r.GET("/classes", Serve(hermes.GET, "class", db))
	r.POST("/classes", Serve(hermes.POST, "class", db))
	r.DELETE("/classes", Serve(hermes.DELETE, "class", db))

	r.GET("/events", Serve(hermes.GET, "event", db))
	r.POST("/events", Serve(hermes.POST, "event", db))
	r.DELETE("/events", Serve(hermes.DELETE, "event", db))

	r.GET("/fixtures", Serve(hermes.GET, "fixture", db))
	r.POST("/fixtures", Serve(hermes.POST, "fixture", db))
	r.DELETE("/fixtures", Serve(hermes.DELETE, "fixture", db))

	r.GET("/links", Serve(hermes.GET, "link", db))
	r.POST("/links", Serve(hermes.POST, "link", db))
	r.DELETE("/links", Serve(hermes.DELETE, "link", db))

	r.GET("/objects", Serve(hermes.GET, "object", db))
	r.POST("/objects", Serve(hermes.POST, "object", db))
	r.DELETE("/objects", Serve(hermes.DELETE, "object", db))

	r.GET("/ontologies", Serve(hermes.GET, "ontology", db))
	r.POST("/ontologies", Serve(hermes.POST, "ontology", db))
	r.DELETE("/ontologies", Serve(hermes.DELETE, "ontology", db))

	r.GET("/relations", Serve(hermes.GET, "relation", db))
	r.POST("/relations", Serve(hermes.POST, "relation", db))
	r.DELETE("/relations", Serve(hermes.DELETE, "relation", db))

	r.GET("/routines", Serve(hermes.GET, "routine", db))
	r.POST("/routines", Serve(hermes.POST, "routine", db))
	r.DELETE("/routines", Serve(hermes.DELETE, "routine", db))

	r.GET("/schedules", Serve(hermes.GET, "schedule", db))
	r.POST("/schedules", Serve(hermes.POST, "schedule", db))
	r.DELETE("/schedules", Serve(hermes.DELETE, "schedule", db))

	r.GET("/tasks", Serve(hermes.GET, "task", db))
	r.POST("/tasks", Serve(hermes.POST, "task", db))
	r.DELETE("/tasks", Serve(hermes.DELETE, "task", db))

	r.GET("/traits", Serve(hermes.GET, "trait", db))
	r.POST("/traits", Serve(hermes.POST, "trait", db))
	r.DELETE("/traits", Serve(hermes.DELETE, "trait", db))

	r.GET("/users", Serve(hermes.GET, "user", db))
	r.POST("/users", Serve(hermes.POST, "user", db))
	r.DELETE("/users", Serve(hermes.DELETE, "user", db))

	return r
}
