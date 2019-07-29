package plug

import (
	"fmt"

	"github.com/lmuench/api/client"
)

func init() {
	Register(Log{})
}

type Log struct{}

func (_ Log) OnReq(req *client.Request) {
	fmt.Printf("plug.Log#OnReq(): %s %s\n", req.Method, req.URL)
}

func (_ Log) OnRes(res *client.Response) {
	fmt.Println("not yet implemented")
}
