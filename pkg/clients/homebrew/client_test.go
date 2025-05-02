package homebrew_test

import (
	"context"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/glamorousis/distillery/pkg/clients/homebrew"
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

func TestHomebrewClient(t *testing.T) {
	tests := []struct {
		name       string
		testFunc   func() error
		shouldFail bool
	}{
		{
			name: "GetFormula_Success",
			testFunc: func() error {
				mockResponse := loadTestData(t, "get-formula.json")
				client := homebrew.NewClient(newMockClient(mockResponse, http.StatusOK))
				_, err := client.GetFormula(context.Background(), "formula-name")
				return err
			},
			shouldFail: false,
		},
		{
			name: "GetFormula_InvalidJSON",
			testFunc: func() error {
				client := homebrew.NewClient(newMockClient("invalid json", http.StatusOK))
				_, err := client.GetFormula(context.Background(), "formula-name")
				return err
			},
			shouldFail: true,
		},
		{
			name: "GetFormula_HTTPError",
			testFunc: func() error {
				client := homebrew.NewClient(newMockClient("", http.StatusInternalServerError))
				_, err := client.GetFormula(context.Background(), "formula-name")
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
