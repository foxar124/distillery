package hashicorp_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/glamorousis/distillery/pkg/clients/hashicorp"
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

func TestHashicorpClient(t *testing.T) {
	client := hashicorp.NewClient(newMockClient("", http.StatusOK))

	tests := []struct {
		name       string
		testFunc   func() error
		shouldFail bool
	}{
		{
			name: "ListProducts_Success",
			testFunc: func() error {
				mockResponse := loadTestData(t, "list-products.json")
				client := hashicorp.NewClient(newMockClient(mockResponse, http.StatusOK))
				_, err := client.ListProducts(context.Background())
				return err
			},
			shouldFail: false,
		},
		{
			name: "ListReleases_Success",
			testFunc: func() error {
				mockResponse := loadTestData(t, "list-releases.json")
				client := hashicorp.NewClient(newMockClient(mockResponse, http.StatusOK))
				_, err := client.ListReleases(context.Background(), "owner/repo", nil)
				return err
			},
			shouldFail: false,
		},
		{
			name: "GetVersion_Success",
			testFunc: func() error {
				mockResponse := loadTestData(t, "get-release.json")
				client := hashicorp.NewClient(newMockClient(mockResponse, http.StatusOK))
				_, err := client.GetVersion(context.Background(), "owner/repo", "v1.0.0")
				return err
			},
			shouldFail: false,
		},
		{
			name: "ListProducts_InvalidURL",
			testFunc: func() error {
				_, err := client.ListProducts(context.Background())
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListReleases_InvalidURL",
			testFunc: func() error {
				_, err := client.ListReleases(context.Background(), "invalid-url-%%", nil)
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetVersion_InvalidURL",
			testFunc: func() error {
				_, err := client.GetVersion(context.Background(), "invalid-url-%%", "v1.0.0")
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListProducts_HTTPError",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.ListProducts(context.Background())
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListReleases_HTTPError",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.ListReleases(context.Background(), "owner/repo", nil)
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetVersion_HTTPError",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.GetVersion(context.Background(), "owner/repo", "v1.0.0")
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListProducts_InvalidJSON",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.ListProducts(context.Background())
				return err
			},
			shouldFail: true,
		},
		{
			name: "ListReleases_InvalidJSON",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.ListReleases(context.Background(), "owner/repo", nil)
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetVersion_InvalidJSON",
			testFunc: func() error {
				client := hashicorp.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.GetVersion(context.Background(), "owner/repo", "v1.0.0")
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
