package http_call

import "net/http"

func GetCaller[Request any, Response any](url string) *HttpCaller[Request, Response] {
	return &HttpCaller[Request, Response]{
		URL:     url,
		Method:  http.MethodGet,
		Headers: make(map[string]any),
		Params:  make(map[string]any),
		Request: nil,
	}
}

func PostCaller[Request any, Response any](url string) *HttpCaller[Request, Response] {
	return &HttpCaller[Request, Response]{
		URL:     url,
		Method:  http.MethodPost,
		Headers: make(map[string]any),
		Params:  make(map[string]any),
	}
}

func PutCaller[Request any, Response any](url string) *HttpCaller[Request, Response] {
	return &HttpCaller[Request, Response]{
		URL:     url,
		Method:  http.MethodPut,
		Headers: make(map[string]any),
		Params:  make(map[string]any),
	}
}

func DeleteCaller[Request any, Response any](url string) *HttpCaller[Request, Response] {
	return &HttpCaller[Request, Response]{
		URL:     url,
		Method:  http.MethodDelete,
		Headers: make(map[string]any),
		Params:  make(map[string]any),
	}
}

type HttpCaller[Request any, Response any] struct {
	URL     string
	Method  string
	Headers map[string]any
	Params  map[string]any
	Request *Request
}

func (c *HttpCaller[Request, Response]) AddHeader(key string, value any) {
	c.Headers[key] = value
}

func (c *HttpCaller[Request, Response]) AddParam(key string, value any) {
	c.Params[key] = value
}

func (c *HttpCaller[Request, Response]) SetRequest(request *Request) {
	c.Request = request
}
