package hermes

import (
	"github.com/elos/data"
	"github.com/elos/data/transfer"
)

func Serve(e *transfer.Envelope, db data.DB) {
	var err error

	switch e.Action {
	case GET:
		err = Get(e, db)
	case POST:
		err = Post(e, db)
	case DELETE:
		err = Delete(e, db)
	case ECHO:
		err = Echo(e, db)
	case SYNC:
		err = Sync(e, db)
	default:
		e.WriteJSON(ErrGeneric)
		return
	}

	// no error
	if err == nil {
		return
	}

	switch err {
	case data.ErrAccessDenial:
		e.WriteJSON(ErrUnauthorized)
	case data.ErrNotFound:
		e.WriteJSON(ErrNotFound)
	default:
		e.WriteJSON(ErrGeneric)
	}
}
