package routes

import (
	"encoding/json"

	"github.com/elos/api/middleware"
	"github.com/elos/api/services"
	"github.com/elos/data/transfer"
	"github.com/elos/ehttp/serve"
	"github.com/elos/models"
)

const ScheduleIDParam = "schedule_id"

func retrieveSchedule(c *serve.Conn, db services.DB) (*models.Schedule, bool) {
	id, ok := retrieveIDParam(ScheduleIDParam, c, db)
	if !ok {
		return nil, false
	}

	schedule, err := models.FindSchedule(db, *id)
	if err != nil {
		ServerError(c, err)
		return nil, false
	}

	return schedule, true
}

func SchedulesGET(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	schedule, ok := retrieveSchedule(c, db)
	if !ok {
		return
	}

	if !checkReadAccess(user, schedule, c, db) {
		return
	}

	c.Response(
		200,
		transfer.StringMap(transfer.Map(schedule)),
	)
}

func SchedulesPOST(c *serve.Conn, db services.DB) {
	user, ok := middleware.RetrieveUser(c, ServerError)
	if !ok {
		return
	}

	// --- Decode the request body {{{
	decoder := json.NewDecoder(c.Request().Body)
	schedule := models.NewSchedule()
	if err := decoder.Decode(schedule); err != nil {
		ServerError(c, err)
		return
	}
	// --- }}}

	// --- Update or Save {{{
	creation := false

	if schedule.Id == "" {
		schedule.SetID(db.NewID())
		creation = true
	}

	if schedule.OwnerID == "" {
		schedule.OwnerID = user.Id
	}

	if user.ID().String() != schedule.OwnerID {
		Unauthorized(c)
		return
	}

	if err := db.Save(schedule); err != nil {
		ServerError(c, err)
		return
	}

	var status int
	if creation {
		status = 201
	} else {
		status = 200
	}
	// --- }}}

	c.Response(
		status,
		transfer.StringMap(transfer.Map(schedule)),
	)
}

func SchedulesDELETE(c *serve.Conn, db services.DB) {}

func SchedulesOPTIONS(c *serve.Conn) {
	c.WriteHeader(200)
}
