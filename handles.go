package api

import (
	"encoding/json"

	"github.com/elos/api/hermes"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/ehttp/sock"
)

func Serve(a transfer.Action, k data.Kind, db data.DB) serve.Route {
	return func(c *serve.Conn) {
		r := c.Request()
		decoder := json.NewDecoder(r.Body)
		data := make(map[data.Kind]data.AttrMap)
		err := decoder.Decode(&data)
		if err != nil {
			panic("hmm")
		}
		e := transfer.NewEnvelope(c, a, data)
		hermes.Serve(e, db)
	}
}

func WebSocket(u sock.Upgrader, man autonomous.Manager) serve.Route {
	return func(c *serve.Conn) {
	}
}
