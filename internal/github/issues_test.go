package github

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestIssuesService_Get(t *testing.T) {
	setupTest()

	mux.Handle("/repos/testOwner/testRepo/issues/1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		testHeader(t, r, testHeaderAccept, testDefaultMediaType)
		testHeader(t, r, testHeaderAPIVersion, testDefaultAPIVersion)
		testHeader(t, r, testHeaderAuthorization, "Bearer "+testToken)

		if r.ContentLength != 0 {
			t.Errorf("Issues.Get() got = %v, want %v", r.ContentLength, 0)
		}

		// create test response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"number":1, "title": "Issue", "body": "Body"}`)
	}))

	issue, resp, err := client.Issues.Get("testOwner", "testRepo", 1)

	// check issue
	want := &Issue{
		Number: Int(1),
		Title:  String("Issue"),
		Body:   String("Body"),
	}
	if !cmp.Equal(issue, want) {
		t.Errorf("Issues.Get() got = %v, want %v", issue, want)
	}

	// check response
	if resp == nil {
		t.Errorf("Issues.Get() got = %v, want %v", issue, want)
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Issues.Get() got = %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// check error
	if err != nil {
		t.Errorf("Issues.Get() error = %v, wantErr %v", err, nil)
		return
	}
}

func TestIssuesService_Update(t *testing.T) {
	setupTest()

	mux.Handle("/repos/testOwner/testRepo/issues/1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPatch)
		testHeader(t, r, testHeaderAccept, testDefaultMediaType)
		testHeader(t, r, testHeaderAPIVersion, testDefaultAPIVersion)
		testHeader(t, r, testHeaderAuthorization, "Bearer "+testToken)

		// create test response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"number":1, "title": "Updated Issue", "body": "Updated Body"}`)
	}))

	// Call the Update method with valid input
	issueRequest := &IssueRequest{
		Title: String("Updated Issue"),
		Body:  String("Updated Body"),
	}
	issue, resp, err := client.Issues.Update("testOwner", "testRepo", 1, issueRequest)

	// check issue
	want := &Issue{
		Number: Int(1),
		Title:  String("Updated Issue"),
		Body:   String("Updated Body"),
	}
	if !cmp.Equal(issue, want) {
		t.Errorf("Issues.Update() got = %v, want %v", issue, want)
	}

	// check response
	if resp == nil {
		t.Errorf("Issues.Update() got = %v, want %v", issue, want)
	}
	if resp != nil && resp.StatusCode != http.StatusOK {
		t.Errorf("Issues.Update() got = %v, want %v", resp.StatusCode, http.StatusOK)
	}

	// check error
	if err != nil {
		t.Errorf("Issues.Update() error = %v, wantErr %v", err, nil)
		return
	}
}

func TestIssuesService_Create(t *testing.T) {
	setupTest()

	// Prepare the issue request
	issueRequest := &IssueRequest{
		Title: String("Test Issue"),
		Body:  String("This is a test issue"),
	}

	mux.Handle("/repos/testOwner/testRepo/issues", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := new(IssueRequest)
		assertNilError(t, json.NewDecoder(r.Body).Decode(v))

		testMethod(t, r, http.MethodPost)
		testHeader(t, r, testHeaderAccept, testDefaultMediaType)
		testHeader(t, r, testHeaderAPIVersion, testDefaultAPIVersion)
		testHeader(t, r, testHeaderAuthorization, "Bearer "+testToken)

		if !cmp.Equal(v, issueRequest) {
			t.Errorf("Issues.Create() got = %v, want %v", v, issueRequest)
		}

		// create test response
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"number":1, "title": "Test Issue", "body": "This is a test issue"}`)
	}))

	// Create the issue
	issue, resp, err := client.Issues.Create("testOwner", "testRepo", issueRequest)

	// check issue
	want := &Issue{
		Number: Int(1),
		Title:  String("Test Issue"),
		Body:   String("This is a test issue"),
	}
	if !cmp.Equal(issue, want) {
		t.Errorf("Issues.Create() got = %v, want %v", issue, want)
	}

	// check response
	if resp == nil {
		t.Errorf("Issues.Create() got = %v, want %v", issue, want)
	}
	if resp != nil && resp.StatusCode != http.StatusCreated {
		t.Errorf("Issues.Create() got = %v, want %v", resp.StatusCode, http.StatusCreated)
	}

	// check error
	if err != nil {
		t.Errorf("Issues.Create() error = %v, wantErr %v", err, nil)
		return
	}
}
