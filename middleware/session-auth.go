package middleware

import "github.com/elos/ehttp/serve"

const UserArtifact = "session-user"

type SessionAuth struct{}

func (sa *SessionAuth) Inbound(c *serve.Conn) bool {
	return true
}

func (sa *SessionAuth) Outbound(c *serve.Conn) bool {
	return true
}
