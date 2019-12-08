package hermes

import (
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/models"
)

const (
	GET    transfer.Action = "GET"
	POST   transfer.Action = "POST"
	DELETE transfer.Action = "DELETE"
	ECHO   transfer.Action = "ECHO"
	SYNC   transfer.Action = "SYNC"
)

func Get(e *transfer.Envelope, db data.DB) error {
	for kind, attrs := range e.Data {
		m := models.ModelFor(kind)
		r, _ := transfer.UnmarshalAttrs(attrs, m)

		if err := db.PopulateByID(m); err != nil {
			return err
		}

		p := transfer.NewPackage(POST, transfer.Map(r))
		if err := e.WriteJSON(p); err != nil {
			return err
		}
	}
	return nil
}

func Post(e *transfer.Envelope, db data.DB) error {
	for kind, attrs := range e.Data {
		m := models.ModelFor(kind)
		r, err := transfer.UnmarshalAttrs(attrs, m)
		if err != nil {
			return err
		}

		if err := db.Save(r); err != nil {
			return err
		}

		p := transfer.NewPackage(POST, transfer.Map(m))
		if err := e.WriteJSON(p); err != nil {
			return err
		}
	}
	return nil
}

func Delete(e *transfer.Envelope, db data.DB) error {
	for kind, attrs := range e.Data {
		m := models.ModelFor(kind)
		r, err := transfer.UnmarshalAttrs(attrs, m)
		if err != nil {
			return err
		}

		if err := db.Delete(r); err != nil {
			return err
		}

		p := transfer.NewPackage(DELETE, transfer.Map(m))
		if err := e.WriteJSON(p); err != nil {
			return err
		}
	}
	return nil
}

func Echo(e *transfer.Envelope, db data.DB) error {
	return e.WriteJSON(e)
}

func Sync(e *transfer.Envelope, db data.DB) error {
	return nil
}
