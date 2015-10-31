package main

import (
	"log"

	"github.com/elos/api"
	apimiddleware "github.com/elos/api/middleware"
	"github.com/elos/api/routes"
	"github.com/elos/autonomous"
	emiddleware "github.com/elos/ehttp/middleware"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func main() {
	// Create and start primary autonomous hub
	hub := autonomous.NewHub()
	go hub.Start()
	hub.WaitStart()

	// Mongo should be running on consul
	// Start mongo
	// mongo.Runner.ConfigFile = "mongo.conf"
	// go hub.StartAgent(mongo.Runner)

	db, err := models.MongoDB("172.16.1.78:27017")
	//db, err := models.MongoDB("mongodb.service.consul:27017")
	if err != nil {
		log.Fatal(err)
	}

	// Setup Middleware
	middleware := &api.Middleware{
		Cors: apimiddleware.NewCors(apimiddleware.AuthHeader),
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
