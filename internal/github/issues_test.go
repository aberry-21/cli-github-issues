package github

import (
	"net/http"
	"reflect"
	"testing"
)

func TestIssuesService_Create(t *testing.T) {
	type args struct {
		owner string
		repo  string
		issue *IssueRequest
	}
	tests := []struct {
		name           string
		args           args
		wantIssueBody  string
		wantIssueTitle string
		wantIssueState string
		wantErr        bool
	}{
		{
			name: "Successfully result",
			args: args{
				owner: testOwner,
				repo:  testRepo,
				issue: &IssueRequest{
					Title: String("test title"),
					Body:  String("test body"),
				},
			},
			wantIssueBody:  "test body",
			wantIssueTitle: "test title",
			wantIssueState: "open",
		},
	}

	c, _ := NewClient(http.DefaultClient, testToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := c.Issues.Create(tt.args.owner, tt.args.repo, tt.args.issue)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check title
			if got.Title != nil && !reflect.DeepEqual(*got.Title, tt.wantIssueTitle) {
				t.Errorf("Create() got.Title = %v, want.Title %v", got, tt.wantIssueTitle)
			}
			// Check issue body
			if got.Body != nil && !reflect.DeepEqual(*got.Body, tt.wantIssueBody) {
				t.Errorf("Create() got.Body = %v, want.Body %v", got, tt.wantIssueBody)
			}
			// Check issue state
			if got.State != nil && !reflect.DeepEqual(*got.State, tt.wantIssueState) {
				t.Errorf("Create() got.State = %v, want.State %v", got, tt.wantIssueState)
			}
		})
	}
}

func TestIssuesService_Get(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		number int
	}
	tests := []struct {
		name           string
		s              IssuesService
		args           args
		wantIssueBody  string
		wantIssueTitle string
		wantIssueState string
		wantErr        bool
	}{
		{
			name: "Successfully result",
			args: args{
				owner:  testOwner,
				repo:   testRepo,
				number: 1,
			},
			wantIssueBody:  "Test",
			wantIssueTitle: "Test",
			wantIssueState: "open",
		},
	}

	c, _ := NewClient(http.DefaultClient, testToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := c.Issues.Get(tt.args.owner, tt.args.repo, tt.args.number)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check title
			if got.Title != nil && !reflect.DeepEqual(*got.Title, tt.wantIssueTitle) {
				t.Errorf("Get() got.Title = %v, want.Title %v", got, tt.wantIssueTitle)
			}
			// Check issue body
			if got.Body != nil && !reflect.DeepEqual(*got.Body, tt.wantIssueBody) {
				t.Errorf("Get() got.Body = %v, want.Body %v", got, tt.wantIssueBody)
			}
			// Check issue state
			if got.State != nil && !reflect.DeepEqual(*got.State, tt.wantIssueState) {
				t.Errorf("Get() got.State = %v, want.State %v", got, tt.wantIssueState)
			}
		})
	}
}

func TestIssuesService_Update(t *testing.T) {
	type args struct {
		owner  string
		repo   string
		number int
		issue  *IssueRequest
	}
	tests := []struct {
		name           string
		args           args
		wantIssueBody  string
		wantIssueTitle string
		wantIssueState string
		wantErr        bool
	}{
		{
			name: "Successfully result",
			args: args{
				owner: testOwner,
				repo:  testRepo,
				issue: &IssueRequest{
					Title: String("test title"),
					Body:  String("test body"),
				},
				number: 1,
			},
			wantIssueBody:  "test body",
			wantIssueTitle: "test title",
			wantIssueState: "open",
		},
	}

	c, _ := NewClient(http.DefaultClient, testToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := c.Issues.Update(tt.args.owner, tt.args.repo, tt.args.number, tt.args.issue)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Check title
			if got.Title != nil && !reflect.DeepEqual(*got.Title, tt.wantIssueTitle) {
				t.Errorf("Update() got.Title = %v, want.Title %v", got, tt.wantIssueTitle)
			}
			// Check issue body
			if got.Body != nil && !reflect.DeepEqual(*got.Body, tt.wantIssueBody) {
				t.Errorf("Update() got.Body = %v, want.Body %v", got, tt.wantIssueBody)
			}
			// Check issue state
			if got.State != nil && !reflect.DeepEqual(*got.State, tt.wantIssueState) {
				t.Errorf("Update() got.State = %v, want.State %v", got, tt.wantIssueState)
			}
		})
	}
}
