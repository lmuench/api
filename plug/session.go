package plug

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/lmuench/api/client"
)

func init() {
	Register(Session{})
}

type Session struct{}

func (_ Session) OnReq(req *client.Request) {
	b, err := ioutil.ReadFile("cookie")
	if err != nil {
		fmt.Println(err)
	}
	req.Cookie = string(b)
}

func (_ Session) OnRes(res *client.Response) {
	b := []byte(
		strings.Join(res.Cookies, "; "),
	)
	err := ioutil.WriteFile("cookie", b, 0600)
	if err != nil {
		fmt.Println(err)
	}
}
