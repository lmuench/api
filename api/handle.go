package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/lmuench/api/plug"
)

type Call struct {
	Command string
	Args    []string
}

type Answer struct {
	Result   string
	Response *http.Response
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
	req, err := http.NewRequest(call.Command, call.Args[0], nil)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range plug.Registry {
		p.OnReq(req)
	}

	fmt.Printf("%s %s\n", call.Command, call.Args[0])
	fmt.Printf("Request cookies: %s\n\n", req.Cookies())

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	for _, p := range plug.Registry {
		p.OnRes(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return Answer{
		Result:   string(body),
		Response: res,
	}
}

func handleDefault(call Call) Answer {
	return Answer{
		Result:   "not implemented",
		Response: nil,
	}
}
