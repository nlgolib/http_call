package http_call

type Response[Data any] struct {
	HttpStatus  int
	RawResponse string
	Data        *Data
}
