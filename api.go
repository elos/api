package api

import (
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

	return &Api{
		router:     router,
		Middleware: m,
		Services:   s,
	}
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context.ClearHandler(http.HandlerFunc(api.router.ServeHTTP)).ServeHTTP(w, r)
}
