package api_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/elos/api"
	"github.com/elos/api/middleware"
	"github.com/elos/api/routes"
	"github.com/elos/autonomous"
	"github.com/elos/data"
	emiddleware "github.com/elos/ehttp/middleware"
	"github.com/elos/models"
)

// --- Setup {{{

var (
	db               data.DB
	middlewareStruct *api.Middleware
	servicesStruct   *api.Services
	apiHandler       *api.Api
	server           *httptest.Server
)

func init() {
	var err error
	db, err = models.MongoDB("localhost")

	if err != nil {
		log.Fatalf("Connection to database failed: %s", err)
	}

	middlewareStruct := &api.Middleware{
		Cors: new(middleware.Cors),
		Log:  new(emiddleware.Null),
		SessionAuth: &middleware.SessionAuth{
			DB:                  db,
			UnauthorizedHandler: routes.Unauthorized,
		},
	}

	servicesStruct := &api.Services{
		Agents: autonomous.NewHub(),
		DB:     db,
	}

	apiHandler := api.New(
		middlewareStruct,
		servicesStruct,
	)

	server = httptest.NewServer(apiHandler)
}

// --- }}}

// --- Describe: /sessions {{{

// --- Factories {{{

func buildUserAndCredential(db data.DB) (*models.User, *models.Credential) {
	user := models.NewUser()
	user.SetID(db.NewID())
	credential := models.NewCredential()
	credential.SetID(db.NewID())

	user.IncludeCredential(credential)
	credential.SetOwner(user)

	credential.Public = models.RandomString(32)
	credential.Private = models.RandomString(32)
	credential.Spec = "password"

	if err := db.Save(user); err != nil {
		log.Fatal(err)
	}

	if err := db.Save(credential); err != nil {
		log.Fatal(err)
	}

	return user, credential
}

// --- }}}

// --- Describe: "GET" {{{

// --- Context: "Valid Request" {{{

func TestSessionsGETValidRequest(t *testing.T) {
	// --- GIVEN: user with a 'password' credential and a session {{{

	_, credential := buildUserAndCredential(db)
	session, err := credential.NewSession(db, 3600*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// --- }}}

	// --- WHEN: GET /sessions session_id query param and appropriate auth header{{{
	u := server.URL + fmt.Sprintf("/sessions?session_id=%s", session.ID())
	request, err := http.NewRequest("GET", u, strings.NewReader(""))
	request.Header.Add(middleware.AuthHeader, session.Token)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err) // something wrong while sending request
	}
	// --- }}}

	// --- THEN: 200 with session {{{
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		t.Error(err)
	}

	// It: should return a status of 200
	if data["status"].(float64) != 200 {
		t.Fatalf("Expected status to be 200, but got %d", data["status"].(float64))
	}

	// It: should return a session
	if data["data"].(map[string]interface{})["session"] == nil {
		t.Fatalf("Expected data to have a session key")
	}
	// --- }}}
}

// --- }}}

// --- Context: "Unauthorized" {{{

func TestSessionsGETUnauthorized(t *testing.T) {
	// --- GIVEN: user with a 'password' credential and a session {{{

	_, credential := buildUserAndCredential(db)
	session, err := credential.NewSession(db, 3600*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	// --- }}}

	// --- WHEN: GET /sessions session_id query param without appropriate auth header{{{
	u := server.URL + fmt.Sprintf("/sessions?session_id=%s", session.ID())
	request, err := http.NewRequest("GET", u, strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err) // something wrong while sending request
	}
	// --- }}}

	// --- THEN: 403 without session {{{
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		t.Error(err)
	}

	// It: should return a status of 403
	if data["status"].(float64) != 403 {
		t.Fatalf("Expected status to be 403, but got %d", data["status"].(float64))
	}
	// --- }}}
}

// --- }}}

// --- }}}

// --- Describe: "POST" {{{

// --- Context: "Valid Request" {{{

func TestSessionsPOSTValidRequest(t *testing.T) {
	// --- GIVEN: A user with one 'password' credential {{{

	user, credential := buildUserAndCredential(db)

	// --- }}}

	// --- WHEN: POST /sessions with public, private and user_id query params {{{

	u := server.URL + fmt.Sprintf("/sessions?public=%s&private=%s&user_id=%s", credential.Public, credential.Private, user.ID())
	request, err := http.NewRequest("POST", u, strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		t.Fatal(err) // something wrong while sending request
	}

	// --- }}}

	// --- THEN: 201 with session {{{
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	t.Log(string(body))

	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		t.Error(err)
	}

	// It: should return a status of 201
	if data["status"].(float64) != 201 {
		t.Fatalf("Expected status to be 201, but got %d", data["status"].(float64))
	}

	// It: should return a new session with token
	if data["data"].(map[string]interface{})["session"] == nil {
		t.Fatalf("Expected data to have a session key")
	}
	// --- }}}
}

// --- }}}

// --- }}}

// --- }}}
