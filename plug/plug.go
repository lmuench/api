package plug

import "github.com/lmuench/api/client"

type Plug interface {
	OnReq(req *client.Request)
	OnRes(res *client.Response)
}
