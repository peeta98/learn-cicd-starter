package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	// Testing Scenarios
	// 1) Authorization headers don't exist
	// 2) Malformed authorization header
	// 3) Happy path, returns apiKey
	tests := map[string]struct{
		input http.Header
		want string
		err error
	}{
		"Unexistent Authorization header": {input: http.Header{}, want: "", err: ErrNoAuthHeaderIncluded},
		"Malformed Authorization header": {input: http.Header{"Authorization": []string{"malformed secret"}}, want:"", err: errors.New("malformed authorization header")},
		"Returns ApiKey": {input: http.Header{"Authorization": []string{"ApiKey secret"}}, want: "secret", err: nil},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)
			if tc.want != got || (err != nil && err.Error() != tc.err.Error()) {
				t.Errorf("expected: %v and %v, got: %v and %v", tc.want, tc.err, got, err)
			}
		})
	}
}