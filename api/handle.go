package api

import (
	"fmt"
	"strings"

	"github.com/lmuench/api/client"
	"github.com/lmuench/api/plug"
)

type Call struct {
	Command string
	Args    []string
}

type Answer struct {
	Result string
}

func Handle(call Call) Answer {
	switch call.Command {
	case "POST", "GET", "PUT", "DELETE":
		return handleRequest(call)
	default:
		return handleDefault(call)
	}
}

func handleRequest(call Call) Answer {
	req := &client.Request{
		Method: call.Command,
		URL:    strings.Join(call.Args, " "),
	}

	for _, p := range plug.Registry {
		p.OnReq(req)
	}

	fmt.Printf("%s %s\n", req.Method, req.URL)
	fmt.Printf("Cookie: %s\n\n", req.Cookie)

	res := client.Fetch(req)

	for _, p := range plug.Registry {
		p.OnRes(res)
	}

	return Answer{
		Result: res.Body,
	}
}

func handleDefault(call Call) Answer {
	return Answer{
		Result: "not implemented",
	}
}
