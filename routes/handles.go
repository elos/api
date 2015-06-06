package routes

import (
	"encoding/json"
	"fmt"

	"github.com/elos/api/hermes"
	"github.com/elos/api/services"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/sock"
	"github.com/elos/models"
)

func Unauthorized(c *serve.Conn) {
	c.Error(403, 4030, "Unauthorized", "Make sure you are providing your access token")
}

func RecordNotFound(c *serve.Conn) {
	c.Error(404, 4040, "Not Found", "Make sure you have a valid id")
}

func ServerError(c *serve.Conn, err error) {
	c.Error(500, 5000, "Server Error", err.Error())
}

func BadParam(c *serve.Conn, param string) {
	c.Error(400, 4000, "Bad param", fmt.Sprintf("Bad parameter: %s", param))
}

func retrieveIDParam(name string, c *serve.Conn, db services.DB) (*data.ID, bool) {
	id, err := db.ParseID(c.ParamVal(name))
	if err != nil {
		BadParam(c, name)
		return new(data.ID), false
	}

	return &id, true
}

func checkReadAccess(user *models.User, property models.Property, c *serve.Conn, db data.DB) bool {
	if canRead, err := user.CanRead(db, property); err != nil {
		ServerError(c, err)
		return false
	} else {
		if !canRead {
			Unauthorized(c)
			return false
		}
	}

	return true
}

func checkWriteAccess(user *models.User, property models.Property, c *serve.Conn, db data.DB) bool {
	if canWrite, err := user.CanWrite(db, property); err != nil {
		ServerError(c, err)
		return false
	} else {
		if !canWrite {
			Unauthorized(c)
			return false
		}
	}

	return true
}

func Serve(a transfer.Action, k data.Kind, db data.DB) serve.Route {
	return func(c *serve.Conn) {
		r := c.Request()
		decoder := json.NewDecoder(r.Body)
		data := make(map[data.Kind]data.AttrMap)
		err := decoder.Decode(&data)
		if err != nil {
			c.WriteJSON(hermes.ErrMalformedData)
			return
		}
		e := transfer.NewEnvelope(c, a, data)
		hermes.Serve(e, db)
	}
}

func WebSocket(u sock.Upgrader, man autonomous.Manager) serve.Route {
	return func(c *serve.Conn) {
	}
}
