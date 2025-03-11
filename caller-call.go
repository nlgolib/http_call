package http_call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *HttpCaller[RequestType, ResponseType]) Call() (*Response[ResponseType], error) {
	url := c.handleParams()
	if c.Debug {
		fmt.Println("[DEBUG] URL:", url)
	}

	var requestBody io.Reader
	if c.Request != nil {
		json, err := json.Marshal(c.Request)
		if err != nil {
			if c.Debug {
				fmt.Println("[DEBUG] Error marshalling request:", err)
			}
			return nil, err
		}
		requestBody = bytes.NewReader(json)
	}

	req, err := http.NewRequest(c.Method, url, requestBody)
	if err != nil {
		if c.Debug {
			fmt.Println("[DEBUG] Error creating request:", err)
		}
		return nil, err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if c.Debug {
			fmt.Println("[DEBUG] Error sending request:", err)
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if c.Debug {
			fmt.Println("[DEBUG] Error reading response:", err)
		}
		return nil, err
	}

	if c.Debug {
		fmt.Println("[DEBUG] Response:", string(body))
	}

	var responseBody ResponseType
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		if c.Debug {
			fmt.Println("[DEBUG] Error unmarshalling response:", err)
		}
		return nil, err
	}

	return &Response[ResponseType]{
		HttpStatus:  resp.StatusCode,
		RawResponse: string(body),
		Data:        &responseBody,
	}, nil
}

func (c *HttpCaller[Request, Response]) handleParams() (url string) {
	if c.Params == nil {
		return c.URL
	}
	return URLParam(c.URL, c.Params)
}
