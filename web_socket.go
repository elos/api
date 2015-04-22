package api

import (
	"github.com/elos/agents"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	"github.com/elos/ehttp/handles"
	"github.com/elos/transfer"
)

func WebSocket(u transfer.WebSocketUpgrader, connMan autonomous.Manager) handles.AccessHandle {
	return handles.WebSocketUpgrade(u, func(conn transfer.SocketConnection, a data.Access) {
		agent := agents.NewClientDataAgent(conn, a)
		connMan.StartAgent(agent)
	})
}

func REPL(u transfer.WebSocketUpgrader, connMan autonomous.Manager) handles.AccessHandle {
	return handles.WebSocketUpgrade(u, func(conn transfer.SocketConnection, a data.Access) {
		agent := agents.NewREPLAgent(conn, a)
		connMan.StartAgent(agent)
	})
}
