package github

import (
	"fmt"
	"net/http"
	"time"
)

type IssuesService service

type Issue struct {
	Number    *int       `json:"number,omitempty"`
	HTMLURL   *string    `json:"html_url,omitempty"`
	Title     *string    `json:"title,omitempty"`
	State     *string    `json:"state,omitempty"`
	User      *User      `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	Body      *string    `json:"body,omitempty"`
}

type User struct {
	Login   *string `json:"login,omitempty"`
	HTMLURL *string `json:"html_url,omitempty"`
}

type IssueRequest struct {
	Title     *string   `json:"title,omitempty"`
	Body      *string   `json:"body,omitempty"`
	Assignees *[]string `json:"assignees,omitempty"`
	Milestone *int      `json:"milestone,omitempty"`
	Labels    *[]string `json:"labels,omitempty"`
	Assignee  *string   `json:"assignee,omitempty"`
	State     *string   `json:"state,omitempty"`
}

// Create a new issue on the specified repository.
//
// GITHUB-API docs: https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#create-an-issue
//
//meta:operation POST /repos/{owner}/{repo}/issues
func (s *IssuesService) Create(owner string, repo string, issue *IssueRequest) (*Issue, *http.Response, error) {
	const op = "github.issue.create"

	// prepare create issue request
	request, err := s.client.NewRequest(
		http.MethodPost,
		fmt.Sprintf("/repos/%s/%s/issues", owner, repo),
		issue,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	// do create issue
	res := new(Issue)
	resp, err := s.client.Do(request, res)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, resp, nil
}

// Get an issue from the specified repository.
//
// GITHUB-API docs: https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#get-an-issue
//
//meta:operation GET /repos/{owner}/{repo}/issues/{issue_number}
func (s *IssuesService) Get(owner string, repo string, number int) (*Issue, *http.Response, error) {
	const op = "github.issue.get"

	// prepare get issue request
	request, err := s.client.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/repos/%s/%s/issues/%d", owner, repo, number),
		nil,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	// do get issue
	res := new(Issue)
	resp, err := s.client.Do(request, res)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, resp, nil
}

// Update an issue on the specified repository.
//
// GITHUB-API docs: https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#update-an-issue
//
//meta:operation PATCH /repos/{owner}/{repo}/issues/{issue_number}
func (s *IssuesService) Update(owner string, repo string, number int, editedIssue *IssueRequest) (*Issue, *http.Response, error) {
	const op = "github.issue.update"

	// prepare update issue request
	request, err := s.client.NewRequest(
		http.MethodPatch,
		fmt.Sprintf("/repos/%s/%s/issues/%d", owner, repo, number),
		editedIssue,
	)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	// do update issue
	res := new(Issue)
	resp, err := s.client.Do(request, res)
	if err != nil {
		return nil, nil, fmt.Errorf("%s: %w", op, err)
	}

	return res, resp, nil
}
