package services

import (
	"github.com/elos/autonomous"
	"github.com/elos/data"
)

// The file contains the interface definition for elos services
// There is only this one file because write these interfaces
// are intermediaries to external services

// DB is an intermediary for the data.dB type
type DB interface {
	data.DB
}

// Agents is an intermediary for the autonomous.Manager type
type Agents interface {
	autonomous.Manager
}
