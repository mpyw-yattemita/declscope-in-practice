package github

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/github/github-mcp-server/pkg/raw"
	"github.com/google/go-github/v73/github"
	"github.com/shurcooL/githubv4"
	"github.com/stretchr/testify/assert"
)

func stubGetClientFn(client *github.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return client, nil
	}
}

func stubGetClientFromHTTPFn(client *http.Client) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return github.NewClient(client), nil
	}
}

func stubGetClientFnErr(err string) GetClientFn {
	return func(_ context.Context) (*github.Client, error) {
		return nil, errors.New(err)
	}
}

func stubGetGQLClientFn(client *githubv4.Client) GetGQLClientFn {
	return func(_ context.Context) (*githubv4.Client, error) {
		return client, nil
	}
}

func stubGetRawClientFn(client *raw.Client) raw.GetRawClientFn {
	return func(_ context.Context) (*raw.Client, error) {
		return client, nil
	}
}

func badRequestHandler(msg string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		structuredErrorResponse := github.ErrorResponse{
			Message: msg,
		}

		b, err := json.Marshal(structuredErrorResponse)
		if err != nil {
			http.Error(w, "failed to marshal error response", http.StatusInternalServerError)
		}

		http.Error(w, string(b), http.StatusBadRequest)
	}
}

func Test_IsAcceptedError(t *testing.T) {
	tests := []struct {
		name           string
		err            error
		expectAccepted bool
	}{
		{
			name:           "github AcceptedError",
			err:            &github.AcceptedError{},
			expectAccepted: true,
		},
		{
			name:           "regular error",
			err:            fmt.Errorf("some other error"),
			expectAccepted: false,
		},
		{
			name:           "nil error",
			err:            nil,
			expectAccepted: false,
		},
		{
			name:           "wrapped AcceptedError",
			err:            fmt.Errorf("wrapped: %w", &github.AcceptedError{}),
			expectAccepted: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := isAcceptedError(tc.err)
			assert.Equal(t, tc.expectAccepted, result)
		})
	}
}
