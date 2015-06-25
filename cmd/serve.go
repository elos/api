package main

import (
	"log"

	"github.com/elos/api"
	apimiddleware "github.com/elos/api/middleware"
	"github.com/elos/api/routes"
	"github.com/elos/autonomous"
	"github.com/elos/data/builtin/mongo"
	emiddleware "github.com/elos/ehttp/middleware"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func main() {
	// Create and start primary autonomous hub
	hub := autonomous.NewHub()
	go hub.Start()
	hub.WaitStart()

	// Start mongo
	// mongo.Runner.ConfigFile = "mongo.conf"
	go hub.StartAgent(mongo.Runner)

	db, err := models.MongoDB("localhost")
	if err != nil {
		log.Fatal(err)
	}

	// Setup Middleware
	middleware := &api.Middleware{
		Cors: new(apimiddleware.Cors),
		Log:  emiddleware.LogRequest,
		SessionAuth: &apimiddleware.SessionAuth{
			DB:                  db,
			UnauthorizedHandler: routes.Unauthorized,
		},
	}

	// Setup Services
	services := &api.Services{
		DB:     db,
		Agents: hub, // run on the main hub for now
	}

	api := api.New(middleware, services)

	serveOptions := &serve.Opts{
		Handler: api,
	}

	apiServer := serve.New(serveOptions)

	hub.StartAgent(apiServer)

	go autonomous.HandleIntercept(hub.Stop)
	hub.WaitStop()
}
