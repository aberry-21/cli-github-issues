package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	headerAccept        = "Accept"
	headerAuthorization = "Authorization"
	headerAPIVersion    = "X-GitHub-Api-Version"

	defaultMediaType  = "application/vnd.github+json"
	defaultAPIVersion = "2022-11-28"
	defaultBaseURL    = "https://api.github.com/"
)

type Client struct {
	client  *http.Client
	BaseUrl *url.URL
	common  service
	Issues  *IssuesService
}

type service struct {
	client *Client
}

func NewClient(client *http.Client, token string) (*Client, error) {
	c := &Client{client: client}
	transport := c.client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	c.client.Transport = roundTripperFunc(
		func(req *http.Request) (*http.Response, error) {
			req = req.Clone(req.Context())
			req.Header.Set(headerAuthorization, fmt.Sprintf("Bearer %s", token))
			return transport.RoundTrip(req)
		},
	)
	err := c.initialize()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) initialize() error {
	c.common.client = c

	if c.BaseUrl == nil {
		var err error
		c.BaseUrl, err = url.Parse(defaultBaseURL)
		if err != nil {
			return err
		}
	}

	c.Issues = (*IssuesService)(&c.common)
	return nil
}

func (c *Client) NewRequest(method string, urlStr string, body any) (*http.Request, error) {
	fullUrl, err := c.BaseUrl.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	bodyReader, err := makeRequestBodyIfExist(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fullUrl.String(), bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set(headerAccept, defaultMediaType)
	req.Header.Set(headerAPIVersion, defaultAPIVersion)
	return req, nil
}

func makeRequestBodyIfExist(body any) (io.ReadWriter, error) {
	if body == nil {
		return nil, nil
	}

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(body)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (c *Client) Do(req *http.Request, res any) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Notice: ignore status code
	if err := json.NewDecoder(resp.Body).Decode(res); err != nil {
		return nil, err
	}
	return resp, nil
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

func (fn roundTripperFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return fn(r)
}

func String(v string) *string { return &v }

func Int(v int) *int { return &v }
