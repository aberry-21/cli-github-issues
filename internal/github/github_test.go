package github

import (
	"io"
	"net/http"
	"reflect"
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

func TestClient_NewRequest(t *testing.T) {
	type args struct {
		method string
		urlStr string
		body   any
	}
	type result struct {
		method  string
		url     string
		body    any
		headers map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    result
		wantErr bool
	}{
		{
			name: "Check nil body",
			args: args{
				method: "POST",
				urlStr: "repos/owner/repo/issues",
				body:   nil,
			},
			want: result{
				method: "POST",
				url:    testDefaultBaseURL + "repos/owner/repo/issues",
				body:   nil,
				headers: map[string]string{
					testHeaderAccept:     testDefaultMediaType,
					testHeaderAPIVersion: testDefaultAPIVersion,
				},
			},
		},
		{
			name: "Check not nil body",
			args: args{
				method: "POST",
				urlStr: "repos/owner/repo/issues",
				body:   &User{Login: String("l")},
			},
			want: result{
				method: "POST",
				url:    testDefaultBaseURL + "repos/owner/repo/issues",
				body:   `{"login":"l"}` + "\n",
				headers: map[string]string{
					testHeaderAccept:     testDefaultMediaType,
					testHeaderAPIVersion: testDefaultAPIVersion,
				},
			},
		},
		{
			name: "Invalid url",
			args: args{
				method: "GET",
				urlStr: "your#$%^&*(proper$#$%%^(password",
				body:   &User{Login: String("l")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(http.DefaultClient, "uhapruhwqedenx223")
			got, err := c.NewRequest(tt.args.method, tt.args.urlStr, tt.args.body)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("NewRequest() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			var body any
			if got.Body != nil {
				bodyArr, _ := io.ReadAll(got.Body)
				body = string(bodyArr)
			}
			res := result{
				method: got.Method,
				url:    got.URL.String(),
				body:   body,
				headers: map[string]string{
					headerAccept:     got.Header.Get(headerAccept),
					headerAPIVersion: got.Header.Get(headerAPIVersion),
				},
			}
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("NewRequest() got = %v, want %v", res, tt.want)
			}
		})
	}
}

//func TestClient_DoInvalidRequest(t *testing.T) {
//	c, _ := NewClient(http.DefaultClient, testToken)
//
//	req, _ := c.NewRequest(http.MethodPost, "invalid/url/test", nil)
//	issue := new(Issue)
//	got, err := c.Do(req, issue)
//
//	if err != nil {
//		t.Errorf("Do() error = %v", err)
//	}
//	t.Errorf("Do() error = %v", got)
//}

func String(v string) *string { return &v }
