package plug

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func init() {
	Register(Session{})
}

type Session struct{}

func (_ Session) OnReq(req *http.Request) {
	b, err := ioutil.ReadFile("cookie")
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Cookie", string(b))
}

func (_ Session) OnRes(res *http.Response) {
	var cookies []string
	for _, cookie := range res.Cookies() {
		cookies = append(cookies, cookie.Name+"="+cookie.Value)
	}

	err := ioutil.WriteFile("cookie", []byte(strings.Join(cookies, "; ")), 0600)
	if err != nil {
		fmt.Println(err)
	}
}
