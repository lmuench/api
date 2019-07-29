package api

import (
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
	plugs := []plug.Plug{
		plug.Log{},
	}

	switch call.Command {
	case "POST", "GET", "PUT", "DELETE":
		return handleRequest(call, plugs)
	default:
		return handleDefault(call)
	}
}

func handleRequest(call Call, plugs []plug.Plug) Answer {
	req := &client.Request{
		Method: call.Command,
		URL:    strings.Join(call.Args, " "),
	}

	for _, p := range plugs {
		p.OnReq(req)
	}

	res := client.Fetch(req)
	return Answer{
		Result: res.Body,
	}
}

func handleDefault(call Call) Answer {
	return Answer{
		Result: "not implemented",
	}
}
