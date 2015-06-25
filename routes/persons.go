package routes

import (
	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func PersonsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// find _or_ create functionality
	if c.ParamVal("person_id") == "" {
		person := models.NewPerson()

		if err := db.PopulateByField("owner_id", user.ID().String(), person); err != nil {
			if err == data.ErrNotFound {
				RecordNotFound(c)
			} else {
				ServerError(c, err)
			}
			return
		}

		c.Response(
			200,
			transfer.StringMap(transfer.Map(person)),
		)

		return
	}

	GET(models.PersonKind, c, db)
}

func PersonsPOST(c *serve.Conn, db services.DB) {
	POST(models.PersonKind, c, db)
}

func PersonsDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.PersonKind, c, db)
}

func PersonsOPTIONS(c *serve.Conn) {
	OPTIONS(models.PersonKind, c)
}
