package http_call

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nlgolib/logger"
)

func (c *HttpCaller[RequestType, ResponseType]) Call() (*Response[ResponseType], error) {
	url := c.handleParams()
	logger.Debug(fmt.Sprintf("URL: %s", url))

	var requestBody io.Reader
	if c.Request != nil {
		json, err := json.Marshal(c.Request)
		if err != nil {
			logger.Error(fmt.Sprintf("Error marshalling request: %s", err))
			return nil, err
		}
		logger.Debug(fmt.Sprintf("Request body: %s", string(json)))
		requestBody = bytes.NewReader(json)
	} else {
		logger.Debug("No request body")
	}

	logger.Debug(fmt.Sprintf("Request: %s %s", c.Method, url))

	req, err := http.NewRequest(c.Method, url, requestBody)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating request: %s", err))
		return nil, err
	}

	for k, v := range c.Headers {
		req.Header.Set(k, fmt.Sprintf("%v", v))
		logger.Debug(fmt.Sprintf("Header: %s %s", k, v))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("Error sending request: %s", err))
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error(fmt.Sprintf("Error reading response: %s", err))
		return nil, err
	}

	logger.Debug(fmt.Sprintf("Response: %s", string(body)))

	var responseBody ResponseType
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		logger.Error(fmt.Sprintf("Error unmarshalling response: %s", err))
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
