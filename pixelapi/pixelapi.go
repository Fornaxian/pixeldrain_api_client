package pixelapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// PixelAPI is the Pixeldrain API client
type PixelAPI struct {
	client      *http.Client
	apiEndpoint string
	key         string
	realIP      string
}

// New creates a new Pixeldrain API client to query the Pixeldrain API with
func New(apiEndpoint string) (api PixelAPI) {
	api.client = &http.Client{Timeout: time.Minute * 5}
	api.apiEndpoint = apiEndpoint

	// Pixeldrain uses unix domain sockets on its servers to minimize latency
	// between the web interface daemon and API daemon. Golang does not
	// understand that it needs to dial a unix socket on this case so we create
	// a custom HTTP transport which uses the unix socket instead of TCP
	if strings.HasPrefix(apiEndpoint, "http://unix:") {
		// Get the socket path from the API endpoint
		var sockPath = strings.TrimPrefix(apiEndpoint, "http://unix:")

		// Fake the dialer to use a unix socket instead of TCP
		api.client.Transport = &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", sockPath)
			},
		}

		// Fake a domain name to stop Go's HTTP client from complaining about
		// the domain name. This string will be completely ignored during
		// requests
		api.apiEndpoint = "http://api.sock"
	} else {
		api.client.Transport = http.DefaultTransport
	}

	return api
}

// Login logs a user into the pixeldrain API. The original PixelAPI does not get
// logged in, only the returned PixelAPI
func (p PixelAPI) Login(apiKey string) PixelAPI {
	p.key = apiKey
	return p
}

// RealIP sets the real IP address to use when making API requests
func (p PixelAPI) RealIP(ip string) PixelAPI {
	p.realIP = ip
	return p
}

// Standard response types

// Error is an error returned by the pixeldrain API. If the request failed
// before it could reach the API the error will be on a different type
type Error struct {
	Status     int    `json:"-"` // One of the http.Status types
	Success    bool   `json:"success"`
	StatusCode string `json:"value"`
	Message    string `json:"message"`

	// In case of the multiple_errors code this array will be populated with
	// more errors
	Errors []Error `json:"errors,omitempty"`

	// Metadata regarding the error
	Extra map[string]interface{} `json:"extra,omitempty"`
}

func (e Error) Error() string { return e.StatusCode }

// ErrIsServerError returns true if the error is a server-side error
func ErrIsServerError(err error) bool {
	if apierr, ok := err.(Error); ok && apierr.Status >= 500 {
		return true
	}
	return false
}

// ErrIsClientError returns true if the error is a client-side error
func ErrIsClientError(err error) bool {
	if apierr, ok := err.(Error); ok && apierr.Status >= 400 && apierr.Status < 500 {
		return true
	}
	return false
}

func (p *PixelAPI) do(r *http.Request) (*http.Response, error) {
	if p.key != "" {
		r.SetBasicAuth("", p.key)
	}
	if p.realIP != "" {
		r.Header.Set("X-Real-IP", p.realIP)
	}

	return p.client.Do(r)
}

func (p *PixelAPI) getString(path string) (string, error) {
	req, err := http.NewRequest("GET", p.apiEndpoint+"/"+path, nil)
	if err != nil {
		return "", err
	}
	resp, err := p.do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	return string(bodyBytes), err
}

func (p *PixelAPI) getRaw(path string) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", p.apiEndpoint+"/"+path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

func (p *PixelAPI) jsonRequest(method, path string, target interface{}) error {
	req, err := http.NewRequest(method, p.apiEndpoint+"/"+path, nil)
	if err != nil {
		return err
	}
	resp, err := p.do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return parseJSONResponse(resp, target)
}

func (p *PixelAPI) form(method, url string, vals url.Values, target interface{}) error {
	req, err := http.NewRequest(method, p.apiEndpoint+"/"+url, strings.NewReader(vals.Encode()))
	if err != nil {
		return fmt.Errorf("prepare request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.do(req)
	if err != nil {
		return fmt.Errorf("do request failed: %w", err)
	}

	defer resp.Body.Close()
	return parseJSONResponse(resp, target)
}

func parseJSONResponse(resp *http.Response, target interface{}) (err error) {
	// Test for client side and server side errors
	if resp.StatusCode >= 400 {
		errResp := Error{Status: resp.StatusCode}
		if err = json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("failed to decode json error: %w", err)
		}
		return errResp
	}

	if target == nil {
		return nil
	}

	if err = json.NewDecoder(resp.Body).Decode(target); err != nil {
		return fmt.Errorf("failed to decode json response: %w", err)
	}

	return nil
}
