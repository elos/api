package routes

import (
	"strings"
	"time"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

func DataGET(c *serve.Conn, db services.DB) {
	GET(models.DatumKind, c, db)
}

func DataPOST(c *serve.Conn, db services.DB) {
	POST(models.DatumKind, c, db)
}

func DataDELETE(c *serve.Conn, db services.DB) {
	DELETE(models.DatumKind, c, db)
}

func DataOPTIONS(c *serve.Conn) {
	OPTIONS(models.DatumKind, c)
}

func DataTagsGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	q := db.NewQuery(models.DatumKind)

	iter, err := q.Select(data.AttrMap{"owner_id": user.ID().String()}).Execute()
	if err != nil {
		ServerError(c, err)
		return
	}

	datum := models.NewDatum()
	tagExistence := make(map[string]bool)

	for iter.Next(datum) {
		for i := range datum.Tags {
			tagExistence[datum.Tags[i]] = true
		}

		datum = models.NewDatum()
	}

	if err := iter.Close(); err != nil {
		ServerError(c, err)
		return
	}

	tags := make([]string, 0)
	for k, _ := range tagExistence {
		tags = append(tags, k)
	}

	c.Response(
		200,
		map[string]interface{}{
			"tags": tags,
		},
	)
}

func DataTagsOPTIONS(c *serve.Conn) {
	OPTIONS(models.DatumKind, c)
}

func DataQueryGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Query Params: tags, startTime, endTime {{{
	tags := make([]string, 0)

	if tagParam := c.ParamVal("tags"); tagParam != "" {
		tags = strings.Split(tagParam, ",")
	}

	// Zero
	startTime := *new(time.Time)
	// Present
	endTime := time.Now()

	// Override start time?
	if startTimeParam := c.ParamVal("start_time"); startTimeParam != "" {
		if start, err := time.Parse(time.RFC3339, startTimeParam); err == nil {
			startTime = start
		}
	}

	// Override end time?
	if endTimeParam := c.ParamVal("end_time"); endTimeParam != "" {
		if end, err := time.Parse(time.RFC3339, endTimeParam); err == nil {
			endTime = end
		}
	}
	// --- }}}

	// --- Find Data {{{
	matchData := make([]*models.Datum, 0)
	datum := models.NewDatum()

	// Query
	q := db.NewQuery(models.DatumKind)

	// Iterate
	iter, err := q.Select(data.AttrMap{"owner_id": user.ID().String()}).Execute()
	if err != nil {
		ServerError(c, err)
		return
	}

	// Filter
	for iter.Next(datum) {
		if datum.Match(tags) && datum.Between(startTime, endTime) {
			matchData = append(matchData, datum)
		}
		datum = models.NewDatum()
	}

	if err := iter.Close(); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	c.Response(
		200,
		map[string]interface{}{
			"data": matchData,
		},
	)
}

func DataQueryOPTIONS(c *serve.Conn) {
	OPTIONS(models.DatumKind, c)
}
