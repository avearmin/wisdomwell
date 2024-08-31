package api

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadParameters(t *testing.T) {
	type user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	tests := []struct {
		name         string
		body         string
		expectParams any
		hasError     bool
	}{
		{
			name: "valid request",
			body: `{"name":"armin","email":"arminonsky@foo.bar"}`,
			expectParams: user{
				Name:  "armin",
				Email: "arminonsky@foo.bar",
			},
			hasError: false,
		},
		{
			name:         "invalid request",
			body:         `{"name":}`,
			expectParams: user{},
			hasError:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/foo", strings.NewReader(test.body))

			params := user{}

			err := readParameters(request, &params)
			if isError(err) == test.hasError {
				t.Fatalf("expected error=%t, but got error=%t", test.hasError, isError(err))
			}
			if params != test.expectParams {
				t.Fatalf("expected body=%v,\nbut got body=%v", test.expectParams, params)
			}
		})
	}
}

func isError(err error) bool {
	return err == nil
}
