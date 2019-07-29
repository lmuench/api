package plug

import (
	"fmt"

	"github.com/lmuench/api/client"
)

func Log(req *client.Request) {
	fmt.Printf("plug.Log: %s %s\n", req.Method, req.URL)
}
