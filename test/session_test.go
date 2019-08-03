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

var cookie1 = &http.Cookie{
	Name:  "api_test_session_cookie",
	Value: "abc123",
}

var cookie2 = &http.Cookie{
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

	http.SetCookie(w, cookie1)
}

func protectedEndpointHandler(w http.ResponseWriter, r *http.Request) {
	cookies := r.Cookies()
	if len(cookies) < 1 {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}

	cookie := cookies[0]
	if cookie.Name != cookie1.Name || cookie.Value != cookie1.Value {
		http.Error(w, "", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, cookie2)
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
	protectedEndpointAnswer := api.Handle(protectedEndpointCall)

	cookie2, err := protectedEndpointAnswer.Response.Request.Cookie(cookie1.Name)
	if err != nil {
		t.Error(err)
	}
	if cookie2.Value != cookie1.Value {
		t.Error("Second request did not contain cookie returned by first response")
	}

	_ = store.Delete("cookie")
}
