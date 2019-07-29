package client

type Request struct {
	Method string
	URL    string
}

type Response struct {
	Body string
}

func Fetch(req Request) Response {
	return Response{
		Body: "client.Fetch(req Request): not yet implemented",
	}
}
