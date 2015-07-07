package api

import (
	"path/filepath"
	"runtime"

	"github.com/elos/api/routes"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

var root string

func init() {
	_, filename, _, _ := runtime.Caller(1)
	root = filepath.Dir(filename)
}

func router(m *Middleware, s *Services) serve.Router {
	router := builtin.NewRouter()

	router.GET(routes.Sessions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.SessionsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Sessions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.SessionsPOST(c, s.DB)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Sessions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.SessionsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Sessions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.SessionsOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Actions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ActionsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Actions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ActionsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Actions, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ActionsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Attributes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.AttributesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Attributes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.AttributesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Attributes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.AttributesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Calendars, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.CalendarsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Calendars, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.CalendarsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Calendars, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.CalendarsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Calendars, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.CalendarsOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Classes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ClassesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Classes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ClassesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Classes, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ClassesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Events, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.EventsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Events, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.EventsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Events, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.EventsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Fixtures, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.FixturesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Fixtures, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.FixturesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Fixtures, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.FixturesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Fixtures, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.FixturesOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Links, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.LinksGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Links, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.LinksPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Links, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.LinksDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Objects, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ObjectsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Objects, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ObjectsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Objects, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.ObjectsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Ontologies, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.OntologiesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Ontologies, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.OntologiesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Ontologies, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.OntologiesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Relations, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RelationsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Relations, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RelationsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Relations, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RelationsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Routines, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RoutinesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Routines, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RoutinesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Routines, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.RoutinesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Schedules, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.SchedulesGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Schedules, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.SchedulesPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Schedules, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.SchedulesDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Schedules, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.SchedulesOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Tasks, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TasksGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Tasks, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TasksPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Tasks, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TasksDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Traits, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TraitsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Traits, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TraitsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Traits, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.TraitsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Users, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.UsersGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Users, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.UsersPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Users, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.UsersDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Users, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.UsersOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Persons, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PersonsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Persons, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PersonsPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Persons, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PersonsDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Persons, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.PersonsOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Data, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.DataGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Data, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.DataPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Data, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.DataDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Data, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.DataOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.DataTags, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.DataTagsGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.DataTags, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.DataTagsOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.DataQuery, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.DataQueryGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.DataQuery, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.DataQueryOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.GET(routes.Property, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PropertyGET(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.POST(routes.Property, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PropertyPOST(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.DELETE(routes.Property, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		if ok := m.SessionAuth.Inbound(c); !ok {
			return
		}

		routes.PropertyDELETE(c, s.DB)

		if ok := m.SessionAuth.Outbound(c); !ok {
			return
		}

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	router.OPTIONS(routes.Property, func(c *serve.Conn) {

		if ok := m.Log.Inbound(c); !ok {
			return
		}

		if ok := m.Cors.Inbound(c); !ok {
			return
		}

		routes.PropertyOPTIONS(c)

		if ok := m.Cors.Outbound(c); !ok {
			return
		}

		if ok := m.Log.Outbound(c); !ok {
			return
		}

	})

	return router
}
