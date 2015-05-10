package api

import (
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
)

type API struct {
	sockets autonomous.Manager
	db      data.DB
	router  serve.Router
}

func New(db data.DB, man autonomous.Manager) *API {
	return &API{
		sockets: man,
		db:      db,
		router:  router(db, man),
	}
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}
