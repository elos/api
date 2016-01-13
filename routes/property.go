package routes

import (
	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

// Exportable domains
var Domains = []data.Kind{
	models.ActionKind,
	models.AttributeKind,
	models.CalendarKind,
	models.EventKind,
	models.FixtureKind,
	models.LinkKind,
	models.ObjectKind,
	models.OntologyKind,
	models.PersonKind,
	models.RelationKind,
	models.RoutineKind,
	models.ScheduleKind,
	models.TaskKind,
	models.TraitKind,
}

// --- PropertyGET {{{
func PropertyGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	property := make(map[data.Kind][]data.Record)

	for _, domain := range Domains {
		property[domain] = make([]data.Record, 0)

		q := db.Query(domain)

		q.Select(data.AttrMap{
			"owner_id": user.ID().String(),
		})

		iter, err := q.Execute()
		if err != nil {
			ServerError(c, err)
		}

		r := models.ModelFor(domain)
		for iter.Next(r) {
			property[domain] = append(property[domain], r)
			r = models.ModelFor(domain)
		}

		if err := iter.Close(); err != nil {
			ServerError(c, err)
		}
	}

	// convert to JSON basically
	interfaceProperty := make(map[string]interface{})
	for k, d := range property {
		interfaceProperty[string(k)] = d
	}

	c.Response(
		200,
		interfaceProperty,
	)
}

// --- }}}

func PropertyPOST(c *serve.Conn, db services.DB)   {}
func PropertyDELETE(c *serve.Conn, db services.DB) {}

func PropertyOPTIONS(c *serve.Conn) {
	c.WriteHeader(200)
}
