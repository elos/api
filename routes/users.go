package routes

import (
	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const UserIDParam = "user_id"

func retrieveUser(c *serve.Conn, db services.DB) (*models.User, bool) {
	id, ok := retrieveIDParam(UserIDParam, c, db)
	if !ok {
		return nil, false
	}

	user, err := models.FindUser(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return user, true
}

func UsersGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	queriedUser, ok := retrieveUser(c, db)
	if !ok {
		return
	}

	if queriedUser.ID().String() != user.ID().String() {
		Unauthorized(c)
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(queriedUser)),
	)
}

func UsersPOST(c *serve.Conn, db services.DB) {
}

func UsersDELETE(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	queriedUser, ok := retrieveUser(c, db)
	if !ok {
		return
	}

	if queriedUser.ID().String() != user.ID().String() {
		Unauthorized(c)
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(queriedUser)),
	)
}

func UsersOPTIONS(c *serve.Conn) {
	OPTIONS(models.UserKind, c)
}
