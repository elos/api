package views

import (
	"fmt"

	"github.com/elos/ehttp/templates"
)

type RoutesContext struct {
}

func (r *RoutesContext) Actions() string {
	return fmt.Sprintf("/actions")
}

func (r *RoutesContext) Attributes() string {
	return fmt.Sprintf("/attributes")
}

func (r *RoutesContext) Calendars() string {
	return fmt.Sprintf("/calendars")
}

func (r *RoutesContext) Classes() string {
	return fmt.Sprintf("/classes")
}

func (r *RoutesContext) Data() string {
	return fmt.Sprintf("/data")
}

func (r *RoutesContext) DataQuery() string {
	return fmt.Sprintf("/data/query")
}

func (r *RoutesContext) DataTags() string {
	return fmt.Sprintf("/data/tags")
}

func (r *RoutesContext) Events() string {
	return fmt.Sprintf("/events")
}

func (r *RoutesContext) Fixtures() string {
	return fmt.Sprintf("/fixtures")
}

func (r *RoutesContext) Links() string {
	return fmt.Sprintf("/links")
}

func (r *RoutesContext) Objects() string {
	return fmt.Sprintf("/objects")
}

func (r *RoutesContext) Ontologies() string {
	return fmt.Sprintf("/ontologies")
}

func (r *RoutesContext) Persons() string {
	return fmt.Sprintf("/persons")
}

func (r *RoutesContext) Property() string {
	return fmt.Sprintf("/property")
}

func (r *RoutesContext) Relations() string {
	return fmt.Sprintf("/relations")
}

func (r *RoutesContext) Routines() string {
	return fmt.Sprintf("/routines")
}

func (r *RoutesContext) Schedules() string {
	return fmt.Sprintf("/schedules")
}

func (r *RoutesContext) Sessions() string {
	return fmt.Sprintf("/sessions")
}

func (r *RoutesContext) Tasks() string {
	return fmt.Sprintf("/tasks")
}

func (r *RoutesContext) Traits() string {
	return fmt.Sprintf("/traits")
}

func (r *RoutesContext) Users() string {
	return fmt.Sprintf("/users")
}

var routesContext = &RoutesContext{}

type context struct {
	Routes *RoutesContext
	Data   interface{}
}

func (c *context) WithData(d interface{}) templates.Context {
	return &context{
		Routes: c.Routes,
		Data:   d,
	}
}

var globalContext = &context{
	Routes: routesContext,
}
