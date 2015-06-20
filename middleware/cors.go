package middleware

import "github.com/elos/ehttp/serve"

type Cors int

func (cors *Cors) Inbound(c *serve.Conn) bool {
	c.Header().Add("Access-Control-Allow-Origin", c.Request().Header.Get("Origin"))
	c.Header().Add("Access-Control-Allow-Credentials", "true")
	c.Header().Add("Access-Control-Allow-Headers", AuthHeader)

	return true
}

func (cors *Cors) Outbound(c *serve.Conn) bool {
	return true
}
