package plug

import (
	"net/http"
)

type Plug interface {
	OnReq(req *http.Request)
	OnRes(res *http.Response)
}
