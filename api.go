package api

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp"
	"github.com/elos/ehttp/handles"
	"github.com/elos/models"
	"github.com/elos/transfer"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	autonomous.Life
	autonomous.Stopper

	*ehttp.Server
	*autonomous.Hub
}

func New(host string, port int, store data.Store) *API {
	hub := autonomous.NewHub()
	server := ehttp.NewServer(host, port, newRouter(hub, store), store)

	return &API{
		Hub:    hub,
		Server: server,

		Life:    autonomous.NewLife(),
		Stopper: make(autonomous.Stopper),
	}
}

func (api *API) Start() {
	go api.Server.Start()
	go api.Hub.Start()

	api.Life.Begin()

	serverstop := make(autonomous.Stopper)
	hubstop := make(autonomous.Stopper)

	go func() {
		api.Server.WaitStop()
		go serverstop.Stop()
	}()

	go func() {
		api.Hub.WaitStop()
		go hubstop.Stop()
	}()

	select {
	case <-serverstop:
		go api.Hub.Stop()
		api.Hub.WaitStop()
	case <-hubstop:
		go api.Server.Stop()
		api.Server.WaitStop()
	case <-api.Stopper:
		go api.Hub.Stop()
		api.Hub.WaitStop()
		go api.Server.Stop()
		api.Server.WaitStop()
	}

	api.Life.End()
}

func newRouter(hub *autonomous.Hub, store data.Store) *httprouter.Router {
	r := httprouter.New()

	r.POST("/v1/users/",
		handles.Access(Post(models.UserKind, Params("name")), data.NewAnonAccess(store)))

	r.POST("/v1/events/",
		handles.Auth(Post(models.EventKind, Params("name")), transfer.Auth(transfer.HTTPCredentialer), store))

	r.GET("/v1/authenticate",
		handles.Auth(WebSocket(transfer.DefaultUpgrader, hub), transfer.Auth(transfer.SocketCredentialer), store))

	r.GET("/v1/repl",
		handles.Auth(REPL(transfer.DefaultUpgrader, hub), transfer.Auth(transfer.SocketCredentialer), store))

	return r
}
