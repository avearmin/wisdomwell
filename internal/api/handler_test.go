package api

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestReadParameters(t *testing.T) {
	tests := []struct {
		name         string
		body         string
		expectParams any
		expectError  error
	}{
		{
			name: "valid request",
			body: `{"name":"armin","email":"arminonsky@foo.bar"}`,
			expectParams: struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}{
				Name:  "armin",
				Email: "arminonsky@foo.bar",
			},
			expectError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "/foo", strings.NewReader(test.body))

			params := struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			}{}

			err := readParameters(request, &params)
			if !errors.Is(err, test.expectError) {
				t.Fatalf("expected error=%v, but got error=%v", test.expectError, err)
			}
			if params != test.expectParams {
				t.Fatalf("expected body=%v,\nbut got body=%v", test.expectParams, params)
			}
		})
	}
}
