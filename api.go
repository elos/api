package api

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	autonomous.Life
	autonomous.Stopper

	server  *serve.Server
	sockets *autonomous.Hub
}

func New(host string, port int, db data.DB) *API {
	hub := autonomous.NewHub()
	server := serve.New(&serve.Opts{
		Handler: router(hub, db),
	})

	return &API{
		Life:    autonomous.NewLife(),
		Stopper: make(autonomous.Stopper),

		server:  server,
		sockets: hub,
	}
}

func (api *API) Start() {
	go api.sockets.Start()
	go api.server.Start()
	api.Life.Begin()

	serverstop := make(autonomous.Stopper)
	hubstop := make(autonomous.Stopper)

	go func() {
		api.server.WaitStop()
		go serverstop.Stop()
	}()

	go func() {
		api.sockets.WaitStop()
		go hubstop.Stop()
	}()

	select {
	case <-serverstop:
		go api.sockets.Stop()
		api.sockets.WaitStop()
	case <-hubstop:
		go api.server.Stop()
		api.server.WaitStop()
	case <-api.Stopper:
		go api.sockets.Stop()
		api.sockets.WaitStop()
		go api.server.Stop()
		api.server.WaitStop()
	}

	api.Life.End()
}

func router(hub *autonomous.Hub, db data.DB) *httprouter.Router {
	r := httprouter.New()
	return r
}
