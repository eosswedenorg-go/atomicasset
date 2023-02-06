package atomicasset

import (
	"context"
	"fmt"
	"strings"

	"github.com/imroc/req/v3"
	"github.com/sonh/qs"
)

// Client interacts with the api
type Client struct {
	URL  string
	Host string
	ctx  context.Context
}

// New Creates a new client object
func New(url string) *Client {
	return NewWithContext(url, nil)
}

func NewWithContext(url string, ctx context.Context) *Client {
	return &Client{
		URL: url,
		ctx: ctx,
	}
}

func isContentType(t string, expected string) bool {
	p := strings.IndexByte(t, ';')
	if p >= 0 {
		t = t[:p]
	}
	return t == expected
}

func (c *Client) send(method string, path string, params interface{}) (*req.Response, error) {
	r := req.C().R()

	if params != nil {
		query, err := qs.NewEncoder().Values(params)
		if err != nil {
			return nil, err
		}
		r.SetQueryString(query.Encode())
	}

	if len(c.Host) > 0 {
		r.SetHeader("Host", c.Host)
	}

	if c.ctx != nil {
		r.SetContext(c.ctx)
	}

	resp, err := r.Send(method, c.URL+path)
	if err != nil {
		return nil, err
	}

	t := resp.GetContentType()
	if !isContentType(t, "application/json") {
		return nil, fmt.Errorf("invalid content-type '%s', expected 'application/json'", t)
	}

	if resp.IsError() {
		apiErr := APIError{}
		if resp.Unmarshal(&apiErr) == nil && apiErr.Success.Valid && !apiErr.Success.Bool {
			return nil, fmt.Errorf("API Error: %s", apiErr.Message.String)
		}
	}

	return resp, err
}

func (c *Client) fetch(method string, url string, params interface{}, resp *APIResponse) (*req.Response, error) {
	r, err := c.send(method, url, params)
	if err == nil {
		// Set HTTPStatusCode
		resp.HTTPStatusCode = r.StatusCode
	}
	return r, err
}
