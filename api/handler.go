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
	if len(call.Args) < 1 {
		return handleDefault(call)
	}
	method := call.Command
	url := call.Args[0]
	switch call.Command {
	case "POST", "PUT", "PATCH":
		if len(call.Args) > 1 {
			body := call.Args[1]
			return handleRequestWithBody(method, url, body)
		}
		return handleRequestWithoutBody(method, url)
	case "GET", "DELETE":
		return handleRequestWithoutBody(method, url)
	default:
		return handleDefault(call)
	}
}

func handleRequestWithBody(method, url, body string) Answer {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%s %s\n", method, url)
	fmt.Printf("Body: %s\n", body)

	res := handleRequest(req)
	answer := handleResponse(res)
	return answer
}

func handleRequestWithoutBody(method, url string) Answer {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("%s %s\n", method, url)

	res := handleRequest(req)
	answer := handleResponse(res)
	return answer
}

func handleRequest(req *http.Request) *http.Response {
	for _, p := range plug.Registry {
		p.OnReq(req)
	}

	fmt.Printf("Request cookies: %s\n", req.Cookies())

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	return res
}

func handleResponse(res *http.Response) Answer {
	fmt.Println(res.Status, "\n")

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
