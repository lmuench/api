package plug

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lmuench/api/store"
)

func init() {
	Register(Session{})
}

type Session struct{}

func (_ Session) OnReq(req *http.Request) {
	cookie, ok := store.Get("cookie")
	if !ok {
		fmt.Println("  - session plug: cookie not found")
		return
	}
	req.Header.Set("Cookie", cookie)
}

func (_ Session) OnRes(res *http.Response) {
	var cookies []string
	for _, cookie := range res.Cookies() {
		cookies = append(cookies, cookie.Name+"="+cookie.Value)
	}

	err := store.Set("cookie", strings.Join(cookies, "; "))
	if err != nil {
		fmt.Println("session plug:", err)
	}
}
