package api

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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
	case "POST", "GET", "PUT", "PATCH", "DELETE":
		req := createRequest(call)
		res := handleRequest(req)
		answer := handleResponse(res)
		return answer
	default:
		return handleDefault(call)
	}
}

func createRequest(call Call) *http.Request {
	if len(call.Args) < 1 {
		log.Panicln("URL parameter missing")
	}

	method := call.Command
	url := call.Args[0]
	fmt.Printf("%s %s\n", method, url)

	var req *http.Request
	var err error

	if len(call.Args) > 1 {
		body := call.Args[1]
		fmt.Printf("Body: %s\n", body)
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Panicln(err)
	}
	return req
}

func handleRequest(req *http.Request) *http.Response {
	runPlugsOnReq(req)

	fmt.Printf("Request cookies: %s\n", req.Cookies())

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func handleResponse(res *http.Response) Answer {
	runPlugsOnRes(res)

	fmt.Printf("%s\n\n", res.Status)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return Answer{
		Result:   string(body),
		Response: res,
	}
}

func runPlugsOnReq(req *http.Request) {
	for _, p := range plug.Registry {
		p.OnReq(req)
	}
}

func runPlugsOnRes(res *http.Response) {
	for _, p := range plug.Registry {
		p.OnRes(res)
	}
}

func handleDefault(call Call) Answer {
	return Answer{
		Result:   "not implemented",
		Response: nil,
	}
}
