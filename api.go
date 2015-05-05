package api

import (
	"net/http"

	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/builtin"
	"github.com/elos/ehttp/serve"
)

type API struct {
	autonomous.Life
	autonomous.Stopper

	sockets *autonomous.Hub
	db      data.DB
	router  serve.Router
}

func New(hub *autonomous.Hub, db data.DB) *API {
	return &API{
		Life:    autonomous.NewLife(),
		Stopper: make(autonomous.Stopper),
		sockets: hub,
		db:      db,
		router:  builtin.NewRouter(),
	}
}

func (api *API) Start() {
	go api.sockets.Start()
	api.Life.Begin()

	hubstop := make(autonomous.Stopper)

	go func() {
		api.sockets.WaitStop()
		go hubstop.Stop()
	}()

	select {
	case <-hubstop:
	case <-api.Stopper:
		go api.sockets.Stop()
		api.sockets.WaitStop()
	}

	api.Life.End()
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.router.ServeHTTP(w, r)
}