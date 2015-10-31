package routes

import (
	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

// --- Actions{GET|POST|DELETE|OPTIONS} {{{

func ActionsGET(c *serve.Conn, db services.DB) {
	GET(models.ActionKind, c, db)
}

func ActionsPOST(c *serve.Conn, db services.DB) {
	POST(models.ActionKind, c, db)
}

func ActionsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ActionKind, c, db)
}

func ActionsOPTIONS(c *serve.Conn) {
	OPTIONS(models.ActionKind, c)
}

// --- }}}

// --- Attributes{GET|POST|DELETE|OPTIONS} {{{

func AttributesGET(c *serve.Conn, db services.DB) {
	GET(models.AttributeKind, c, db)
}

func AttributesPOST(c *serve.Conn, db services.DB) {
	POST(models.AttributeKind, c, db)
}

func AttributesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.AttributeKind, c, db)
}

func AttributesOPTIONS(c *serve.Conn) {
	OPTIONS(models.AttributeKind, c)
}

// --- }}}

// --- Calendars{GET|POST|DELETE|OPTIONS} {{{

func CalendarsGET(c *serve.Conn, db services.DB) {
	GET(models.CalendarKind, c, db)
}

func CalendarsPOST(c *serve.Conn, db services.DB) {
	POST(models.CalendarKind, c, db)
}

func CalendarsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.CalendarKind, c, db)
}

func CalendarsOPTIONS(c *serve.Conn) {
	OPTIONS(models.CalendarKind, c)
}

// --- }}}

// --- Events{GET|POST|DELETE|OPTIONS} {{{

func EventsGET(c *serve.Conn, db services.DB) {
	GET(models.EventKind, c, db)
}

func EventsPOST(c *serve.Conn, db services.DB) {
	POST(models.EventKind, c, db)
}

func EventsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.EventKind, c, db)
}

func EventsOPTIONS(c *serve.Conn) {
	OPTIONS(models.EventKind, c)
}

// --- }}}

// --- Fixtures{GET|POST|DELETE|OPTIONS} {{{

func FixturesGET(c *serve.Conn, db services.DB) {
	GET(models.FixtureKind, c, db)
}

func FixturesPOST(c *serve.Conn, db services.DB) {
	POST(models.FixtureKind, c, db)
}

func FixturesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.FixtureKind, c, db)
}

func FixturesOPTIONS(c *serve.Conn) {
	OPTIONS(models.FixtureKind, c)
}

// --- }}}

// --- Links{GET|POST|DELETE|OPTIONS} {{{

func LinksGET(c *serve.Conn, db services.DB) {
	GET(models.LinkKind, c, db)
}

func LinksPOST(c *serve.Conn, db services.DB) {
	POST(models.LinkKind, c, db)
}

func LinksDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.LinkKind, c, db)
}

func LinksOPTIONS(c *serve.Conn) {
	OPTIONS(models.LinkKind, c)
}

// --- }}}

// --- Objects{GET|POST|DELETE|OPTIONS} {{{

func ObjectsGET(c *serve.Conn, db services.DB) {
	GET(models.ObjectKind, c, db)
}

func ObjectsPOST(c *serve.Conn, db services.DB) {
	POST(models.ObjectKind, c, db)
}

func ObjectsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.ObjectKind, c, db)
}

func ObjectsOPTIONS(c *serve.Conn) {
	OPTIONS(models.ObjectKind, c)
}

// --- }}}

// --- Ontologies{GET|POST|DELETE|OPTIONS} {{{

func OntologiesGET(c *serve.Conn, db services.DB) {
	GET(models.OntologyKind, c, db)
}

func OntologiesPOST(c *serve.Conn, db services.DB) {
	POST(models.OntologyKind, c, db)
}

func OntologiesDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.OntologyKind, c, db)
}

func OntologiesOPTIONS(c *serve.Conn) {
	OPTIONS(models.OntologyKind, c)
}

// --- }}}

// --- Groups{GET|POST|DELETE|OPTIONS} {{{

func GroupsGET(c *serve.Conn, db services.DB) {
	GET(models.GroupKind, c, db)
}

func GroupsPOST(c *serve.Conn, db services.DB) {
	POST(models.GroupKind, c, db)
}

func GroupsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.GroupKind, c, db)
}

func GroupsOPTIONS(c *serve.Conn) {
	OPTIONS(models.GroupKind, c)
}

// --- }}}
