package gitlab_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/glamorousis/distillery/pkg/clients/gitlab"
)

func loadTestData(t *testing.T, filename string) string {
	data, err := os.ReadFile("testdata/" + filename)
	assert.NoError(t, err)
	return string(data)
}

func newMockClient(responseBody string, statusCode int) *http.Client {
	return &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) *http.Response {
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(responseBody)),
				Header:     make(http.Header),
			}
		}),
	}
}

type roundTripperFunc func(req *http.Request) *http.Response

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func TestListReleases(t *testing.T) {
	mockResponse := loadTestData(t, "list-releases.json")
	client := gitlab.NewClient(newMockClient(mockResponse, http.StatusOK))
	client.SetToken("test-token")

	releases, err := client.ListReleases(context.Background(), "owner/repo")
	assert.NoError(t, err)
	assert.NotNil(t, releases)
	assert.Equal(t, 2, len(releases))
}

func TestGetLatestRelease(t *testing.T) {
	mockResponse := loadTestData(t, "latest-release.json")
	client := gitlab.NewClient(newMockClient(mockResponse, http.StatusOK))
	client.SetToken("test-token")

	release, err := client.GetLatestRelease(context.Background(), "owner/repo")
	assert.NoError(t, err)
	assert.NotNil(t, release)
	assert.Equal(t, "v1.0.0", release.TagName)
}

func TestGetRelease(t *testing.T) {
	mockResponse := loadTestData(t, "get-release.json")
	client := gitlab.NewClient(newMockClient(mockResponse, http.StatusOK))
	client.SetToken("test-token")

	release, err := client.GetRelease(context.Background(), "owner/repo", "v1.0.0")
	assert.NoError(t, err)
	assert.NotNil(t, release)
	assert.Equal(t, "v1.0.0", release.TagName)
}

func TestGitlabClientErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		testFunc   func() error
		shouldFail bool
	}{
		{
			name: "ListReleases_InvalidURL",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusOK))
				client.SetToken("test-token")
				_, err := client.ListReleases(context.Background(), "invalid-url-%%")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetLatestRelease_InvalidURL",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusOK))
				client.SetToken("test-token")
				_, err := client.GetLatestRelease(context.Background(), "invalid-url-%%")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetRelease_InvalidURL",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusOK))
				client.SetToken("test-token")
				_, err := client.GetRelease(context.Background(), "invalid-url-%%", "v1.0.0")
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListReleases_HTTPError",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.ListReleases(context.Background(), "owner/repo")
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListReleases_InvalidJSON",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.ListReleases(context.Background(), "owner/repo")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetLatestRelease_HTTPError",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.GetLatestRelease(context.Background(), "owner/repo")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetLatestRelease_InvalidJSON",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.GetLatestRelease(context.Background(), "owner/repo")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetRelease_HTTPError",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.GetRelease(context.Background(), "owner/repo", "v1.0.0")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetRelease_InvalidJSON",
			testFunc: func() error {
				client := gitlab.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.GetRelease(context.Background(), "owner/repo", "v1.0.0")
				return err
			},
			shouldFail: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := tc.testFunc()
			if tc.shouldFail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
