package http_call

import (
	"fmt"
	"strings"
)

func URLParam(url string, params map[string]any) string {
	if len(params) == 0 {
		return url
	}
	url = strings.TrimRight(url, "/")
	url += "?"
	for k, v := range params {
		url += fmt.Sprintf("%s=%v&", k, v)
	}
	return url
}
