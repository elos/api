package main

import (
	"log"

	"github.com/elos/autonomous"
	"github.com/elos/data/builtin/mongo"
	"github.com/elos/models"
)

func main() {
	hub := autonomous.NewHub()
	go hub.Start()
	hub.WaitStart()

	go hub.StartAgent(mongo.Runner)

	db, err := models.MongoDB("localhost")
	if err != nil {
		log.Fatal(err)
	}

	user := models.NewUser()
	user.SetID(db.NewID())

	credential := models.NewCredential()
	credential.SetID(db.NewID())
	credential.Public = "public"
	credential.Private = "private"
	credential.Spec = "password"

	user.IncludeCredential(credential)
	credential.SetOwner(user)

	person := models.NewPerson()
	person.SetID(db.NewID())
	person.SetOwner(user)

	if err := db.Save(user); err != nil {
		log.Fatal(err)
	}

	if err := db.Save(credential); err != nil {
		log.Fatal(err)
	}

	if err := db.Save(person); err != nil {
		log.Fatal(err)
	}
}
