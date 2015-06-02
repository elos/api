package routes

import "testing"

// --- Describe: GET "/sessions" {{{

// --- Context: Valid Request {{{

func TestSessionsGETValid(t *testing.T) {
	u := &session.
	NewSessionForUser
	w := httptest.NewRecorder()
	c := serve.NewConn(
		w,
		&http.Request{},
		TestParams{
			"session_id":
		}
	)
}

// --- Context: Valid Request }}}

// --- Describe: GET "/sessions" }}}

// --- Describe: POST "/sessions" {{{

// --- Context: Valid Request {{{

func TestSessionsPOSTValid(t *testing.T) {
}

// --- Context: Valid Request }}}

// --- Describe: POST "/sessions" }}}
