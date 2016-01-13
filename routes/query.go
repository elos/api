package routes

import (
	"encoding/json"
	"log"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
	"github.com/elos/models/access"
)

type query struct {
	Kind               data.Kind
	Space              string
	Skip, Limit, Batch int
	Attrs              data.AttrMap
}

func QueryPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	log.Print("Query Body: ", c.Request().Body)
	decoder := json.NewDecoder(c.Request().Body)
	query := &query{}
	if err := decoder.Decode(query); err != nil {
		ServerError(c, err)
		return
	}

	log.Print("Query: %+v", query)

	q := db.Query(query.Kind)
	q.Select(query.Attrs)
	q.Skip(query.Skip)
	q.Limit(query.Limit)
	q.Batch(query.Batch)

	iter, err := q.Execute()
	if err != nil {
		ServerError(c, err)
		return
	}

	modelList := make([]data.Record, 0)

	model := models.ModelFor(query.Kind)
	for iter.Next(model) {
		log.Printf("Model: %+v", model)
		if canRead, err := access.CanRead(db, access.WrapUser(user), model); err != nil {
			ServerError(c, err)
			return
		} else {
			if canRead {
				modelList = append(modelList, model)
			} else {
				log.Print("Cant read")
			}
		}

		model = models.ModelFor(query.Kind)
	}

	log.Printf("Models to return: %+v", modelList)

	res := make(map[string]interface{})
	res[query.Space] = modelList

	c.Response(
		200,
		res,
	)
}

func QueryOPTIONS(c *serve.Conn) {
	c.WriteHeader(200)
}
