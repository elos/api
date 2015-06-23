package api

import (
	"log"
	"net/http"

	"github.com/elos/api/services"
	"github.com/elos/ehttp/serve"
	"github.com/gorilla/context"
)

type Middleware struct {
	Cors serve.Middleware

	Log serve.Middleware

	SessionAuth serve.Middleware
}

type Services struct {
	services.Agents

	services.DB
}

type Api struct {
	router serve.Router
	*Middleware
	*Services
}

func New(m *Middleware, s *Services) *Api {
	router := router(m, s)

	if m.Cors == nil {
		log.Fatal("Middleware Cors is nil")
	}

	if m.Log == nil {
		log.Fatal("Middleware Log is nil")
	}

	if m.SessionAuth == nil {
		log.Fatal("Middleware SessionAuth is nil")
	}

	if s.Agents == nil {
		log.Fatal("Service Agents is nil")
	}

	if s.DB == nil {
		log.Fatal("Service DB is nil")
	}

	return &Api{
		router:     router,
		Middleware: m,
		Services:   s,
	}
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.ClearHandler(http.HandlerFunc(api.router.ServeHTTP)).ServeHTTP(w, r)
}
