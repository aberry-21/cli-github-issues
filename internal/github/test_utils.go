package github

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	testHeaderAccept        = "Accept"
	testHeaderAuthorization = "Authorization"
	testHeaderAPIVersion    = "X-GitHub-Api-Version"

	testDefaultMediaType  = "application/vnd.github+json"
	testDefaultAPIVersion = "2022-11-28"
	testDefaultBaseURL    = "https://api.github.com/"
	testOwner             = "aberry-21"
	testRepo              = "BeautifulNumber"
	testToken             = "tkn"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	mux *http.ServeMux

	// client is the GitHub client being tested.
	client *Client

	// server is a test HTTP server used to provide mock API responses.
	server *httptest.Server
)

func setupTest() {
	// test server
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	// github client configured to use test server
	client, _ = NewClient(http.DefaultClient, testToken)
	url, _ := url.Parse(server.URL + "/")
	client.BaseUrl = url
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header, want string) {
	t.Helper()
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

func assertNilError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
