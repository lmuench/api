package test

import (
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

func handler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, cookie1)
}

func runServer() {
	http.HandleFunc("/foo", handler)
	log.Fatal(http.ListenAndServe(":6000", nil))
}

func TestSessionTokens(t *testing.T) {
	go runServer()

	call := api.Call{
		Command: "GET",
		Args:    []string{"http://localhost:6000/foo"},
	}
	api.Handle(call)
	answer2 := api.Handle(call)
	cookie2, err := answer2.Response.Request.Cookie(cookie1.Name)
	if err != nil {
		t.Error(err)
	}
	if cookie2.Value != cookie1.Value {
		t.Error("Second request did not contain cookie returned by first response")
	}

	_ = store.Delete("cookie")
}
