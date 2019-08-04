package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	"github.com/lmuench/api/api"
	"github.com/lmuench/api/store"
)

var res1cookie = &http.Cookie{
	Name:  "api_test_session_cookie",
	Value: "abc123",
}

var res2cookie = &http.Cookie{
	Name:  "api_test_session_cookie",
	Value: "def456",
}

type credentials struct {
	Username string
	Password string
}

var validCreds = credentials{
	Username: "bob",
	Password: "pass1",
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	var creds credentials
	err = json.Unmarshal(body, &creds)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if creds != validCreds {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, res1cookie)
}

func protectedEndpointHandler(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if len(cookies) < 1 {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	cookie := cookies[0]
	if cookie.Name != res1cookie.Name || cookie.Value != res1cookie.Value {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, res2cookie)
}

func runServer() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected-endpoint", protectedEndpointHandler)
	log.Fatal(http.ListenAndServe(":6000", nil))
}

func TestSessionTokens(t *testing.T) {
	go runServer()

	loginBody, err := json.Marshal(validCreds)
	if err != nil {
		t.Error(err)
	}

	loginCall := api.Call{
		Command: "POST",
		Args:    []string{"http://localhost:6000/login", string(loginBody)},
	}
	api.Handle(loginCall)

	protectedEndpointCall := api.Call{
		Command: "GET",
		Args:    []string{"http://localhost:6000/protected-endpoint"},
	}
	protectedEndpointAnswer1 := api.Handle(protectedEndpointCall)

	if len(protectedEndpointAnswer1.Response.Request.Cookies()) < 1 {
		t.Error("Second request did not have a cookie in the request header")
	}
	req2cookie := protectedEndpointAnswer1.Response.Request.Cookies()[0]
	if err != nil {
		t.Error(err)
	}
	if req2cookie.Value != res1cookie.Value {
		t.Error("Second request did not contain cookie returned by first response")
	}

	protectedEndpointAnswer2 := api.Handle(protectedEndpointCall)

	if len(protectedEndpointAnswer2.Response.Request.Cookies()) < 1 {
		t.Error("Third request did not have a cookie in the request header")
	}
	req3cookie := protectedEndpointAnswer2.Response.Request.Cookies()[0]
	if err != nil {
		t.Error(err)
	}
	if req3cookie.Value != res2cookie.Value {
		t.Error("Third request did not contain cookie returned by second response")
	}

	// Note: third request will result in 401 Unauthorized but it does not matter for the test

	_ = store.Delete("cookie")
}
