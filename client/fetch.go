package client

type Request struct {
	Method string
	URL    string
	Cookie string
}

type Response struct {
	Body    string
	Cookies []string
}

// mocked
func Fetch(req *Request) *Response {
	return &Response{
		Body:    "Hello, world!",
		Cookies: []string{"foo=qwerty", "bar=asdfg"},
	}
}
