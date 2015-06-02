package services

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
)

type DB interface {
	data.DB
}

type Agents interface {
	autonomous.Manager
}
